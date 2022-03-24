package internal

import (
	"context"
	"regexp"
	"strings"
	"time"

	"github.com/jiuhuche120/jpr/config"
	"github.com/jiuhuche120/jpr/git"
)

const (
	CheckMainBranchMerged   = "checkMainBranchMerged"
	CheckPullRequestTimeout = "checkPullRequestTimeout"
)

func (s *Server) checkMerged(pull git.PullRequest) bool {
	res := s.client.Get(pull.Url + "/merge")
	return strings.Contains(string(res), "No Content")
}

func (s *Server) checkMainBranchMerged(name string, gits config.Gits, rule *config.CheckMainBranchMerged) {
	_, err := s.scheduler.ScheduleWithCron(func(ctx context.Context) {
		s.log.Infof("start check %v main branch merged", name)
		pulls, err := s.client.GetAllPullRequests(gits)
		if err != nil {
			s.log.Error(err)
			return
		}
		var callPulls []git.PullRequest
		for i := 0; i < len(pulls); i++ {
			reg := regexp.MustCompile(rule.Head)
			if pulls[i].State == "open" && reg.FindString(pulls[i].Base.Ref) != "" {
				flag := false
				for j := 0; j < len(pulls); j++ {
					if i == j {
						continue
					}
					if i != j && pulls[i].Title == pulls[j].Title && pulls[j].Base.Ref == rule.Base && pulls[j].State == "open" {
						flag = true
						break
					}
					if i != j && pulls[i].Title == pulls[j].Title && pulls[j].Base.Ref == rule.Base && pulls[j].State == "closed" && s.checkMerged(pulls[j]) {
						flag = true
						break
					}
				}
				if !flag {
					s.log.Infof("the pull request %v from %v to %v is lost", pulls[i].Title, pulls[i].Head.Ref, rule.Base)
					s.setUser(&pulls[i])
					callPulls = append(callPulls, pulls[i])
				}
			}
		}
		if len(callPulls) != 0 {
			s.callHook(callPulls, CheckMainBranchMerged)
		}
		s.log.Infof("stop check %v main branch merged", name)
	}, rule.Cron)
	if err != nil {
		s.log.Infof("task has been scheduled")
	}
}
func (s *Server) checkPullRequestTimeout(name string, gits config.Gits, rule *config.CheckPullRequestTimeout) {
	_, err := s.scheduler.ScheduleWithCron(func(ctx context.Context) {
		s.log.Infof("start check %v pull request timeout", name)
		pulls, err := s.client.GetAllPullRequests(gits)
		if err != nil {
			s.log.Error(err)
			return
		}
		var callPulls []git.PullRequest
		for i := 0; i < len(pulls); i++ {
			if pulls[i].State == "open" {
				timeout, err := time.ParseDuration(rule.Timeout)
				if err != nil {
					s.log.Error(err)
				}
				if time.Since(pulls[i].CreateAt) >= timeout {
					s.log.Infof("the pull request %v is timeout", pulls[i].Title)
					s.setUser(&pulls[i])
					callPulls = append(callPulls, pulls[i])
				}
			}
		}
		if len(callPulls) != 0 {
			s.callHook(callPulls, CheckPullRequestTimeout)
		}
		s.log.Infof("stop check %v pull request timeout", name)
	}, rule.Cron)
	if err != nil {
		s.log.Infof("task has been scheduled")
	}
}
