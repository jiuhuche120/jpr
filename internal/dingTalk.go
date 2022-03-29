package internal

import (
	"sync"

	"github.com/jiuhuche120/jpr/git"
)

var mutex sync.Mutex

func (s *Server) callHook(pull []git.PullRequest, typ string) {
	mutex.Lock()
	defer mutex.Unlock()
	switch typ {
	case CheckMainBranchMerged:
		s.dingTalk.Post(s.config.Webhook, git.NewMsg(pull, "需要合并分支到master分支"))
	case CheckPullRequestTimeout:
		s.dingTalk.Post(s.config.Webhook, git.NewMsg(pull, "PR存活超时"))
	}
}

func (s *Server) setUser(pull *git.PullRequest) {
	pull.DingTalk = s.config.DingTalk[pull.User.Login].Phone
}
