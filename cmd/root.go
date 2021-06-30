package cmd

import (
	"fmt"
	"github.com/urfave/cli"
	"os"
)

func Execute() {
	app := cli.NewApp()
	app.Name = "Sonic"
	app.Description = "This is a set of tools and functions that help accelerate the development of modules."

	app.Action = initCmd.Action

	app.Commands = []cli.Command{
		*initCmd,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

}
