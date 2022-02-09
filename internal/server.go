package internal

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/jiuhuche120/jpr/config"
	"github.com/jiuhuche120/jpr/git"
	"github.com/prometheus/common/log"
)

type Server struct {
	config config.Config
	ctx    context.Context
	cancel context.CancelFunc
}

func NewServer(config config.Config) *Server {
	ctx, cancel := context.WithCancel(context.Background())
	return &Server{config: config, ctx: ctx, cancel: cancel}
}

func (s *Server) GetAllPullRequests() ([]git.PullRequest, error) {
	return s.GetPullRequestByStatus("all")
}

func (s *Server) GetPullRequestByStatus(status string) ([]git.PullRequest, error) {
	client := git.NewClient(
		git.AddHeader("Accept", "application/vnd.github.v3+json"),
		git.AddHeader("Authorization", "token "+s.config.Token),
	)
	return client.GetPullRequests(getUrl(s.config.Owner, s.config.Repo) + "?state=" + status)
}

func (s *Server) GetParentPullRequest(title string) (git.PullRequest, error) {
	pulls, err := s.GetAllPullRequests()
	if err != nil {
		return git.PullRequest{}, err
	}
	for _, pull := range pulls {
		if pull.Title == title {
			return pull, nil
		}
	}
	return git.PullRequest{}, nil
}

func (s *Server) Start() {
	log.Infof("start server")
	go func() {
		ticker := time.NewTicker(s.config.Time)
		defer ticker.Stop()
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-ticker.C:
				log.Infof("start check pull request")
				pulls, err := s.GetAllPullRequests()
				if err != nil {
					log.Error(err)
					return
				}
				for i := 0; i < len(pulls); i++ {
					reg := regexp.MustCompile(s.config.Head)
					var callPulls []git.PullRequest
					if pulls[i].State == "open" && reg.FindString(pulls[i].Base.Ref) != "" {
						for j := 0; j < len(pulls); j++ {
							if i == j {
								continue
							}
							if i != j && pulls[i].Title == pulls[j].Title && pulls[j].Base.Ref == s.config.Base && pulls[j].State == "open" {
								break
							}
							if i != j && pulls[i].Title == pulls[j].Title && pulls[j].Base.Ref == s.config.Base && pulls[j].State == "close" && s.checkMerged(pulls[j]) {
								break
							}
							log.Infof("the pull request from %v to %v is lost", pulls[i].Head.Ref, s.config.Base)
							s.getUser(&pulls[i])
							callPulls = append(callPulls, pulls[i])
						}
						s.callHook(callPulls)
					}
				}
				log.Infof("stop check pull request")
			}
		}
	}()
}

func (s *Server) Stop() {
	s.cancel()
	log.Infof("stop server")
}

func (s *Server) callHook(pull []git.PullRequest) {
	msg := git.NewMsg(pull)
	client := git.NewClient(
		git.AddHeader("Content-Type", "application/json"),
	)
	client.Post(s.config.WebHook, msg)
}
func (s *Server) checkMerged(pull git.PullRequest) bool {
	client := git.NewClient(
		git.AddHeader("Accept", "application/vnd.github.v3+json"),
		git.AddHeader("Authorization", "token "+s.config.Token),
	)
	res := client.Get(pull.Url + "/merge")
	return strings.Contains(string(res), "No Content")
}

func (s *Server) getUser(pull *git.PullRequest) {
	pull.DingTalk = s.config.Users[pull.User.Login]
}

func getUrl(owner, repo string) string {
	return "https://api.github.com/repos/" + owner + "/" + repo + "/pulls"
}
