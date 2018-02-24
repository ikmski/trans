package main

type trans struct {
	input  *input
	screen *screen
}

func newTrans() *trans {

	t := new(trans)

	t.input = newInput()
	t.screen = newScreen(t.input)

	return t
}

func (t *trans) run() {

}
