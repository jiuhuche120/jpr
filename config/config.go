package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

const DefaultPath = "~/.jpr/"
const DefaultName = "config.json"

type Config struct {
	Owner   string            `json:"owner"`
	Repo    string            `json:"repo"`
	Cron    string            `json:"cron"`
	Base    string            `json:"base"`
	Head    string            `json:"head"`
	Token   string            `json:"token"`
	WebHook string            `json:"webhook"`
	Users   map[string]string `json:"users"`
}

func (c *Config) check() error {
	if c.Owner == "" {
		return fmt.Errorf("the config of owner is nil")
	}
	if c.Repo == "" {
		return fmt.Errorf("the config of repo is nil")
	}
	if c.Base == "" {
		return fmt.Errorf("the config of base is nil")
	}
	if c.Head == "" {
		return fmt.Errorf("the config of head is nil")
	}
	if c.Token == "" {
		return fmt.Errorf("the config of token is nil")
	}
	if c.WebHook == "" {
		return fmt.Errorf("the config of webhook is nil")
	}
	if len(c.Users) == 0 {
		return fmt.Errorf("the config of users is nil")
	}
	return nil
}

func LoadConfig() (Config, error) {
	config := Config{}
	path, err := PathRoot()
	if err != nil {
		return Config{}, err
	}
	bytes, err := ioutil.ReadFile(filepath.Join(path, DefaultName))
	if err != nil {
		return Config{}, err
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func PathRoot() (string, error) {
	return homedir.Expand(DefaultPath)
}
