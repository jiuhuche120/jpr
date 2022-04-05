package config

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/mitchellh/go-homedir"
)

const DefaultPath = "~/.jpr/"
const DefaultName = "config.toml"

type Config struct {
	Token    string              `toml:"token"`
	Webhook  string              `toml:"webhook"`
	Gits     map[string]Gits     `toml:"gits"`
	Rules    Rules               `toml:"rules"`
	Log      Log                 `toml:"log"`
	DingTalk map[string]DingTalk `toml:"dingTalk"`
}

type Gits struct {
	Owner string   `toml:"owner"`
	Repo  string   `toml:"repo"`
	Rules []string `toml:"rules"`
}

type Rules struct {
	CheckMainBranchMerged   map[string]*CheckMainBranchMerged   `toml:"checkMainBranchMerged"`
	CheckPullRequestTimeout map[string]*CheckPullRequestTimeout `toml:"checkPullRequestTimeout"`
}

type CheckMainBranchMerged struct {
	Base string `toml:"base"`
	Head string `toml:"head"`
	Cron string `toml:"cron"`
}

type CheckPullRequestTimeout struct {
	Timeout string `toml:"timeout"`
	Cron    string `toml:"cron"`
}

type Log struct {
	Level string `toml:"level"`
}

type DingTalk struct {
	Phone string `toml:"phone"`
	Email string `toml:"email"`
}

func (c *Config) check() error {
	config, err := LoadConfig()
	if err != nil {
		return err
	}
	if config.Token == "" {
		return fmt.Errorf("config token can't empty")
	}
	if config.Webhook == "" {
		return fmt.Errorf("config webhook can't empty")
	}
	if len(config.Gits) == 0 {
		return fmt.Errorf("config gits can't empty")
	}
	for k, v := range config.Gits {
		for _, rule := range v.Rules {
			str := strings.Split(rule, ".")
			switch str[0] {
			case "checkMainBranchMerged":
				if config.Rules.CheckMainBranchMerged[str[1]] == nil {
					return fmt.Errorf("pkg %v's rule %v is not exsit", k, rule)
				}
			case "checkPullRequestTimeout":
				if config.Rules.CheckPullRequestTimeout[str[1]] == nil {
					return fmt.Errorf("pkg %v's rule %v is not exsit", k, rule)
				}
			}
		}
	}
	if len(config.DingTalk) == 0 {
		return fmt.Errorf("config dingTalk can't empty")
	}
	return nil
}

func (c *Config) GetCheckMainBranchMergedRule(gits Gits) *CheckMainBranchMerged {
	for _, rule := range gits.Rules {
		str := strings.Split(rule, ".")
		if str[0] == "checkMainBranchMerged" {
			return c.Rules.CheckMainBranchMerged[str[1]]
		}
	}
	return nil
}

func (c *Config) GetCheckPullRequestTimeoutRule(gits Gits) *CheckPullRequestTimeout {
	for _, rule := range gits.Rules {
		str := strings.Split(rule, ".")
		if str[0] == "checkPullRequestTimeout" {
			return c.Rules.CheckPullRequestTimeout[str[1]]
		}
	}
	return nil
}

func LoadConfig() (*Config, error) {
	path, err := DefaultConfigPath()
	if err != nil {
		return nil, err
	}
	config := Config{}
	if _, err := toml.DecodeFile(path, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func PathRoot() (string, error) {
	return homedir.Expand(DefaultPath)
}

func DefaultConfigPath() (string, error) {
	path, err := homedir.Expand(DefaultPath)
	if err != nil {
		return "", err
	}
	return filepath.Join(path, DefaultName), nil
}
