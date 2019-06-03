package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "tree"

	app.Usage = "list all file"
	app.Action = func(c *cli.Context) error {
		var dir string = "."
		if c.NArg() > 0 {
			dir = c.Args()[0]
		}

		ListDir(dir, 1)
		return nil
	}
}

func ListDir(dirPath string, deep int) (err error) {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return err
	}
	if deep == 1 {
		fmt.Printf("!---%s\n", filepath.Base(dirPath))
	}
	sep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() {
			fmt.Printf("|")
			for i:=0; i<deep; i++ {
				fmt.Printf("    |")
			}
			fmt.Printf("----%s\n", fi.Name())
			ListDir(dirPath + sep + fi.Name(), deep + 1)
			continue
		}

		fmt.Printf("|")
		for i:=0; i<deep; i++ {
			fmt.Printf("    |")
		}
		fmt.Printf("----%s\n", fi.Name())
	}
	return nil
}
