package main

import (
	"fmt"
	"github.com/godaner/sid/command"
	"github.com/urfave/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "sid"
	app.Usage = "New snow id or parse snow id time"

	app.Commands = []cli.Command{
		command.NewSIDCommand,
		command.ParseTimeCommand,
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
