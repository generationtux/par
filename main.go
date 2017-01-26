package main

import (
	"os"
	"path"
	"strings"

	"io/ioutil"

	"fmt"

	"github.com/smallfish/simpleyaml"
	"github.com/urfave/cli"
)

func extractKeys(fileData []byte, ymlPathArgs []string) []string {
	keys := make([]string, 0, 50)
	data := ""
	if len(ymlPathArgs) > 0 {
		y, _ := simpleyaml.NewYaml(fileData)
		data, _ = y.GetPath(ymlPathArgs...).String()
	} else {
		data = string(fileData[:])
	}
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
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Println(file + " Does not exist")
			os.Exit(1)
		}
		fileData, _ := ioutil.ReadFile(file)
		envKeys = append(envKeys, extractKeys(fileData, ymlPathArgs))
	}
	return envKeys
}

func CompareEnvArrays(envKeys [][]string, filenames []string) bool {
	for i := 0; i < len(envKeys); i++ {
		for j := i + 1; j < len(envKeys); j++ {
			if len(envKeys[i]) != len(envKeys[j]) {
				fmt.Println("Unequal number of keys")
				return false
			}
			for y := 0; y < len(envKeys[j]); y++ {
				keyExists := false
				for z := 0; z < len(envKeys[j]); z++ {
					if envKeys[j][y] == envKeys[i][z] {
						keyExists = true
					}
				}
				if !keyExists {
					fmt.Println("Key " + envKeys[j][y] + " doesn't exist in file parameter " + filenames[j])
					return false
				}
			}
		}
	}
	return true
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
		ymlArgs := []string{}
		cmdArgs := c.Args().Get(1)
		if cmdArgs != "" {
			ymlArgs = strings.Split(c.Args().Get(1), ",")
		}
		envKeys := ParseFileData(fileNames, ymlArgs)
		result := CompareEnvArrays(envKeys, fileNames)
		if result {
			fmt.Println("All configs are matching")
			os.Exit(0)
		} else {
			fmt.Println("Keys arent equal")
			os.Exit(1)
		}
		return nil
	}

	app.Run(os.Args)
}
