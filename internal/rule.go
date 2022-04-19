package internal

import (
	"context"
	"encoding/json"
	"regexp"
	"strings"
	"time"

	"github.com/jiuhuche120/jpr/config"
	"github.com/jiuhuche120/jpr/pkg"
)

const (
	CheckMainBranchMerged   = "checkMainBranchMerged"
	CheckPullRequestTimeout = "checkPullRequestTimeout"
)

func (s *Server) checkMerged(pull pkg.PullRequest) bool {
	res, err := s.client.Get(pull.Url + "/merge")
	if err != nil {
		s.log.Error(err)
	}
	return strings.Contains(string(res), "No Content")
}

func (s *Server) IsWorkingDay() bool {
	today := time.Now().Format("2006-01-02")
	bytes, err := s.client.Get("https://timor.tech/api/holiday/info/" + today)
	if err != nil {
		s.log.Error(err)
		return false
	}
	var day pkg.Day
	err = json.Unmarshal(bytes, &day)
	if err != nil {
		s.log.Errorf("holiday api is error %v", err)
		return false
	}
	if day.Type.Type == 0 {
		return true
	}
	return false
}

func (s *Server) checkMainBranchMerged(name string, gits config.Gits, rule *config.CheckMainBranchMerged) {
	_, err := s.scheduler.ScheduleWithCron(func(ctx context.Context) {
		if s.IsWorkingDay() {
			s.log.Infof("start check %v main branch merged", name)
			pulls, err := s.client.GetAllPullRequests(gits)
			if err != nil {
				s.log.Error(err)
				return
			}
			var callPulls []pkg.PullRequest
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
		} else {
			s.log.Infof("today is not working day, skip")
		}
	}, rule.Cron)
	if err != nil {
		s.log.Infof("task has been scheduled")
	}
}
func (s *Server) checkPullRequestTimeout(name string, gits config.Gits, rule *config.CheckPullRequestTimeout) {
	_, err := s.scheduler.ScheduleWithCron(func(ctx context.Context) {
		if s.IsWorkingDay() {
			s.log.Infof("start check %v pull request timeout", name)
			pulls, err := s.client.GetAllPullRequests(gits)
			if err != nil {
				s.log.Error(err)
				return
			}
			var callPulls []pkg.PullRequest
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
		} else {
			s.log.Infof("today is not working day, skip")
		}
	}, rule.Cron)
	if err != nil {
		s.log.Infof("task has been scheduled")
	}
}
