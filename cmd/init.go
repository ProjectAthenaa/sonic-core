package cmd

import (
	"fmt"
	"github.com/otiai10/copy"
	"github.com/urfave/cli"
	"go/build"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
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

		if err := copy.Copy(getSonicCoreDir(gopath), "./"); err != nil {
			return err
		}
		log.Println("Finished generating files!")
		log.Println("Initialising go module")

		modInitCmd := exec.Command("go", "mod", "init", "main")
		modInitCmd.Dir = "./src"
		err := modInitCmd.Run()
		if err != nil {
			return err
		}

		_ = os.Remove("./go.mod")
		_ = os.Remove("./go.main")

		return nil
	},
}

var matchFile = regexp.MustCompile(`sonic-core@(.*)\\(.*)`)

func getSonicCoreDir(gopath string) string {
	var files []string
	filepath.Walk(fmt.Sprintf(`%s\pkg\mod\github.com\!project!athenaa`, gopath), func(path string, info fs.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})

	var versions []string

	for _, file := range files {
		if strings.Contains(file, "sonic-core") {
			matches := matchFile.FindStringSubmatch(file)
			if len(matches) >= 1 {
				versions = append(versions, strings.Split(matches[1], "\\")[0])
			}
		}
	}

	sort.Strings(versions)
	latestVersion := versions[0]
	return fmt.Sprintf("%s\\pkg\\mod\\github.com\\!project!athenaa\\sonic-core@%s", gopath, latestVersion)
}
