package cmd

import (
	"fmt"
	"github.com/otiai10/copy"
	"github.com/urfave/cli"
	"go/build"
	"log"
	"os"
)

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "Generate a new module setup",
	Action: func(ctx *cli.Context) error {
		log.Println("Generating files")
		gopath := os.Getenv("GOPATH")
		if gopath == "" {
			gopath = build.Default.GOPATH
		}
		log.Println(fmt.Sprintf("%s/src/github.com/ProjectAthenaa/sonic/template", gopath))

		info, _ := os.Stat(`C:\Users\sn3ak\go\src\github.com\ProjectAthenaa\sonic\template`)
		log.Println(info)
		if err := copy.Copy(fmt.Sprintf("%s/src/github.com/ProjectAthenaa/sonic/template", gopath), ""); err != nil {
			return err
		}
		log.Println("Finished generating files!")
		return nil
	},
}
