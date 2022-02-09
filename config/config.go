package config

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

const DefaultPath = "~/.jpr/"
const DefaultName = "config.toml"

type Config struct {
	Owner   string            `toml:"owner"`
	Repo    string            `toml:"repo"`
	Time    time.Duration     `toml:"time"`
	Base    string            `toml:"base"`
	Head    string            `toml:"head"`
	Token   string            `toml:"token"`
	WebHook string            `toml:"webhook"`
	Users   map[string]string `toml:"users"`
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
	config := DefaultConfig()
	path, err := PathRoot()
	if err != nil {
		return Config{}, err
	}
	v := viper.New()
	v.SetConfigFile(filepath.Join(path, DefaultName))
	v.SetConfigType("toml")
	v.AutomaticEnv()
	v.SetEnvPrefix("JPR")
	err = v.ReadInConfig()
	if err != nil {
		return Config{}, err
	}
	err = v.Unmarshal(&config)
	if err != nil {
		return Config{}, err
	}
	if err := config.check(); err != nil {
		return Config{}, err
	}
	return config, nil
}

func DefaultConfig() Config {
	return Config{
		Time: time.Hour * 3,
		Head: "release*",
		Base: "master",
	}
}

func PathRoot() (string, error) {
	return homedir.Expand(DefaultPath)
}
