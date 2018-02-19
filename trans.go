package main

import (
	"bufio"
	"os"

	termbox "github.com/nsf/termbox-go"
)

type trans struct {
	input  *input
	screen *screen

	inputCh chan bool
	eventCh chan termbox.Event
}

func newTrans() *trans {

	t := new(trans)

	t.eventCh = make(chan termbox.Event)
	t.inputCh = make(chan bool)

	t.input = newInput()
	t.screen = newScreen(t.input)

	return t
}

func (t *trans) run() {

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	go func() {
		for {
			t.eventCh <- termbox.PollEvent()
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)

	go func() {
		for scanner.Scan() {
			t.input.text = scanner.Text()
			t.inputCh <- true
		}
	}()

	t.screen.render()

mainloop:
	for {

		select {
		case e := <-t.eventCh:

			switch e.Type {

			case termbox.EventKey:

				switch e.Key {

				case termbox.KeyEsc:
					break mainloop

				default:
					t.screen.render()

				}

			default:
				t.screen.render()

			}

		}
	}

}
