package cmd

import (
	"github.com/otiai10/copy"
	"github.com/urfave/cli"
	"log"
)

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Generate a new module setup",
	Action: func(ctx *cli.Context) error{
		log.Println("Generating files")
		if err := copy.Copy("github.com/ProjectAthenaa/template", ""); err != nil {
			log.Fatal(err)
		}
		log.Println("Finished generating files!")
	},
}
