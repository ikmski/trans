package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type screen struct {
}

func newScreen() *screen {

	s := new(screen)

	return s
}

func (s *screen) render(t *translation) {

	s.clear()

	fmt.Print("# Google Translation\n\n")
	fmt.Printf("Text: %s\n\n", t.text)
	fmt.Printf("Result: %s\n\n", t.result)
	fmt.Print("> ")
}

func (s *screen) clear() {

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

}
