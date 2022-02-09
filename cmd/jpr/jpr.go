package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := cli.NewApp()
	app.Name = "Jpr"
	app.Usage = "Jpr is a tool to help merge pull requests"

	app.Commands = []*cli.Command{
		initCmd,
		startCmd,
		versionCmd,
	}
	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
