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

func (s *screen) render(text string) {

	if runtime.GOOS == "windows" {
		cmd := exec.Command("cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}

	fmt.Print("> ")
}
