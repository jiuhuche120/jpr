package main

import (
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"

	"github.com/jiuhuche120/jpr/config"
	"github.com/jiuhuche120/jpr/internal"
	"github.com/urfave/cli/v2"
)

var startCmd = &cli.Command{
	Name:   "start",
	Usage:  "start server",
	Action: Start,
}

func Start(ctx *cli.Context) error {
	path, err := config.PathRoot()
	if err != nil {
		return err
	}
	_, err = os.Stat(filepath.Join(path, config.DefaultName))
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Please initialize config first")
		}
	} else {
		c, err := config.LoadConfig()
		if err != nil {
			return err
		}
		server := internal.NewServer(c)
		var wg sync.WaitGroup
		wg.Add(1)
		handleShutdown(server, &wg)
		server.Start()
		wg.Wait()
	}
	return nil
}

func handleShutdown(server *internal.Server, wg *sync.WaitGroup) {
	var stop = make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM)
	signal.Notify(stop, syscall.SIGINT)
	go func() {
		<-stop
		fmt.Println("received interrupt signal, shutting down...")
		server.Stop()
		wg.Done()
		os.Exit(0)
	}()
}
