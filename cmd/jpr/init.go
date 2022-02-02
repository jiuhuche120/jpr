package main

import (
	"bufio"
	"fmt"
	"github.com/jiuhuche120/jpr/config"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"path/filepath"
)

const DefaultConfig = "config/config.toml"

var initCmd = &cli.Command{
	Name:   "init",
	Usage:  "init config home for jpr",
	Action: Initialize,
}

func Initialize(ctx *cli.Context) error {
	path, err := config.PathRoot()
	if err != nil {
		return err
	}
	_, err = os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			err := os.Mkdir(path, 0777)
			if err != nil {
				return err
			}
		}
	}
	_, err = os.Stat(filepath.Join(path, config.DefaultName))
	if err != nil {
		if os.IsNotExist(err) {
			data, err := ioutil.ReadFile(DefaultConfig)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(path, config.DefaultName), data, 0777)
			if err != nil {
				return err
			}
		}
	}else{
		fmt.Println("jpr configuration file already exists")
		fmt.Println("reinitializing would overwrite your configuration, Y/N?")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		if input.Text() == "Y" || input.Text() == "y"{
			data, err := ioutil.ReadFile(DefaultConfig)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filepath.Join(path, config.DefaultName), data, 0777)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
