package main

import (
	"os"
	"path"
	"strings"

	"fmt"

	"io/ioutil"

	"github.com/smallfish/simpleyaml"
	"github.com/urfave/cli"
)

func extractKeys(fileData []byte, ymlPathArgs []string) []string {
	keys := make([]string, 0, 50)
	y, _ := simpleyaml.NewYaml(fileData)
	data, _ := y.GetPath(ymlPathArgs...).String()
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		keyVals := strings.Split(line, "=")
		if len(keyVals) > 0 {
			if len(keyVals[0]) > 0 {
				keys = append(keys, keyVals[0])
			}
		}
	}
	return keys
}

func ParseFileData(filenames []string, ymlPathArgs []string) [][]string {
	envKeys := make([][]string, 0, 50)
	for _, file := range filenames {
		fileData, _ := ioutil.ReadFile(file)
		envKeys = append(envKeys, extractKeys(fileData, ymlPathArgs))
	}
	return envKeys
}

func appendStringsToCWD(filenames []string) []string {
	pwd, _ := os.Getwd()
	for index, file := range filenames {
		filenames[index] = path.Join(pwd, file)
	}
	return filenames
}

func main() {
	app := cli.NewApp()
	app.Name = "par"
	app.Usage = "fight the loneliness!"

	app.Action = func(c *cli.Context) error {
		fileNames := strings.Split(c.Args().Get(0), ",")
		ymlArgs := strings.Split(c.Args().Get(0), ",")
		envKeys := ParseFileData(fileNames, ymlArgs)
		fmt.Println(envKeys)
		return nil
	}

	app.Run(os.Args)
}
