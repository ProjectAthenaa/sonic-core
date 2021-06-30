package cmd

import (
	"fmt"
	"github.com/otiai10/copy"
	"github.com/urfave/cli"
	"go/build"
	"log"
	"os"
	"os/exec"
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

		if err := copy.Copy(fmt.Sprintf("%s\\src\\github.com\\ProjectAthenaa\\sonic\\template", gopath), "./"); err != nil {
			return err
		}
		log.Println("Finished generating files!")
		log.Println("Initialising go module")

		cdCmd := exec.Command("cd", "./src")
		err := cdCmd.Run()
		if err != nil {
			return err
		}

		modInitCmd := exec.Command("go", "mod init main")
		err = modInitCmd.Run()
		if err != nil {
			return err
		}

		modTidyCmd := exec.Command("go", "mod tidy")
		err = modTidyCmd.Run()
		if err != nil {
			return err
		}
		return nil
	},
}
