package internal

import (
	"strings"

	"github.com/jiuhuche120/jpr/config"
	"github.com/jiuhuche120/jpr/git"
	"github.com/procyon-projects/chrono"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config    config.Config
	scheduler chrono.TaskScheduler
	client    git.Client
	dingTalk  git.Client
	log       logrus.Logger
}

func NewServer(config *config.Config) *Server {
	client := git.NewClient(
		git.AddHeader("Accept", "application/vnd.github.v3+json"),
		git.AddHeader("Authorization", "token "+config.Token),
	)
	dingTalk := git.NewClient(
		git.AddHeader("Content-Type", "application/json"),
	)
	return &Server{config: *config, client: *client, dingTalk: *dingTalk, log: *logrus.New()}
}

func (s *Server) Start() {
	level, err := logrus.ParseLevel(s.config.Log.Level)
	if err != nil {
		s.log.Error(err)
	}
	s.log.SetLevel(level)
	s.log.Infof("start server")
	s.scheduler = chrono.NewDefaultTaskScheduler()
	for name, gits := range s.config.Gits {
		s.log.Infof("start %v check", name)
		for _, rule := range gits.Rules {
			str := strings.Split(rule, ".")
			switch str[0] {
			case CheckMainBranchMerged:
				r := s.config.GetCheckMainBranchMergedRule(gits)
				go s.checkMainBranchMerged(name, gits, r)
			case CheckPullRequestTimeout:
				r := s.config.GetCheckPullRequestTimeoutRule(gits)
				go s.checkPullRequestTimeout(name, gits, r)
			}
		}
	}
}

func (s *Server) Stop() {
	s.scheduler.Shutdown()
	s.log.Infof("stop server")
}
