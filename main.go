package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

var config globalConfig

const (
	configFileName = "config.toml"
)

func mainAction(c *cli.Context) error {

	translation := newTranslation()
	screen := newScreen()

	screen.render(translation)

	doneCh := make(chan bool)
	inputCh := make(chan string)
	defer close(doneCh)
	defer close(inputCh)

	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {

			if scanner.Text() == "quit" {
				doneCh <- true
				break

			} else {

				inputCh <- scanner.Text()
			}
		}
	}()

loop:
	for {
		select {

		case <-doneCh:
			break loop

		case text := <-inputCh:
			translation.do(text)
			screen.render(translation)

		}
	}

	return nil
}

func getConfigFilePath() string {

	filePath := ""
	isExist := false
	curDir, err := os.Getwd()
	if err == nil {
		filePath = filepath.Join(curDir, configFileName)
		_, err = os.Stat(filePath)
		if err == nil {
			isExist = true
		}
	}

	if !isExist {
		filePath = filepath.Join(os.Getenv("HOME"), ".config", "trans", configFileName)
	}

	return filePath
}

func main() {

	configFile := getConfigFilePath()
	_, err := os.Stat(configFile)
	if err != nil {

		config = getDefaultConfig()
		err = config.save(configFile)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	} else {

		_, err := toml.DecodeFile(getConfigFilePath(), &config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	}

	app := cli.NewApp()
	app.Name = "trans"
	app.Usage = "translator"
	app.Description = "command-line translation tool using the Google translation API"
	app.Version = version

	app.Action = mainAction

	app.RunAndExitOnError()
}
