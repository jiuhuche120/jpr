package main

import (
	"fmt"
	"github.com/jiuhuche120/jpr"
	"github.com/urfave/cli/v2"
)

var versionCmd = &cli.Command{
	Name:   "version",
	Usage:  "Jpr version",
	Action: version,
}

func version(ctx *cli.Context) error {
	fmt.Printf("Jpr version: %s-%s-%s\n", jpr.CurrentVersion, jpr.CurrentBranch, jpr.CurrentCommit)
	fmt.Printf("App build date: %s\n", jpr.BuildDate)
	return nil
}
