package main

import (
	"github.com/nsf/termbox-go"
	"github.com/urfave/cli"
)

var (
	version  string
	revision string
)

func drawLine(x, y int, str string) {

	color := termbox.ColorDefault
	backgroundColor := termbox.ColorDefault

	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		termbox.SetCell(x+i, y, runes[i], color, backgroundColor)
	}
}

func draw() {

	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	drawLine(0, 0, "Press ESC to exit.")

	termbox.Flush()
}

func pollEvent() {

	draw()

mainloop:
	for {
		switch ev := termbox.PollEvent(); ev.Type {

		case termbox.EventKey:

			switch ev.Key {

			case termbox.KeyEsc:
				break mainloop

			default:
				draw()

			}

		default:
			draw()

		}
	}

}

func mainAction(c *cli.Context) error {

	// init
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	pollEvent()

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
