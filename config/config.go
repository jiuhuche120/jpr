package config

import (
	"fmt"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"path/filepath"
	"time"
)

const DefaultPath = "~/.jpr/"
const DefaultName = "config.toml"

type Config struct {
	Repo string        `toml:"repo"`
	Time time.Duration `toml:"time"`
	Base string        `toml:"base"`
	Head string        `toml:"head"`
}

func (c *Config) check() error {
	if c.Repo == "" {
		return fmt.Errorf("the config of repo is nil")
	}
	if c.Base == "" {
		return fmt.Errorf("the config of base is nil")
	}
	if c.Head == "" {
		return fmt.Errorf("the config of head is nil")
	}
	return nil
}

func LoadConfig() (*Config, error) {
	var config Config
	path, err := PathRoot()
	if err != nil {
		return nil, err
	}
	v := viper.New()
	v.SetConfigFile(filepath.Join(path, DefaultName))
	v.SetConfigType("toml")
	v.AutomaticEnv()
	v.SetEnvPrefix("JPR")
	err = v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	err = v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func PathRoot() (string, error) {
	return homedir.Expand(DefaultPath)
}
