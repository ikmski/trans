package main

import "github.com/urfave/cli"

var (
	version  string
	revision string
)

func mainAction(c *cli.Context) error {

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
