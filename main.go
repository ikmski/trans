package main

import (
	"bufio"
	"os"

	"github.com/urfave/cli"
)

var (
	version  string
	revision string
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

func main() {

	app := cli.NewApp()
	app.Name = "trans"
	app.Usage = "translator"
	app.Description = "command-line tool for Google translator"
	app.Version = version

	app.Action = mainAction

	app.RunAndExitOnError()
}
