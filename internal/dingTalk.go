package internal

import (
	"sync"

	"github.com/jiuhuche120/jpr/pkg"
)

var mutex sync.Mutex

func (s *Server) callHook(pull []pkg.PullRequest, typ string) {
	mutex.Lock()
	defer mutex.Unlock()
	switch typ {
	case CheckMainBranchMerged:
		s.dingTalk.Post(s.config.Webhook, pkg.NewMsg(pull, "需要合并分支到master分支"))
	case CheckPullRequestTimeout:
		s.dingTalk.Post(s.config.Webhook, pkg.NewMsg(pull, "PR存活超时"))
	}
}

func (s *Server) setUser(pull *pkg.PullRequest) {
	pull.DingTalk = s.config.DingTalk[pull.User.Login].Phone
}
