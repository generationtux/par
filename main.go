package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		fileNames := strings.Split(c.Args().Get(0), ",")
		for _, element := range fileNames {
			fmt.Println(element)
		}
		return nil
	}

	app.Run(os.Args)
}
