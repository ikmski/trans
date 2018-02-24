package main

import (
	"bufio"
	"os"
)

type trans struct {
	screen *screen
}

func newTrans() *trans {

	t := new(trans)

	t.screen = newScreen()

	return t
}

func (t *trans) run() {

	t.screen.render("")

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
			t.screen.render(text)

		}
	}

}
