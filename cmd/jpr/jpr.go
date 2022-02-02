package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Jpr"
	app.Usage = "Jpr is a tool to help merge pull requests"

	app.Commands = []*cli.Command{
		initCmd,
		versionCmd,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
