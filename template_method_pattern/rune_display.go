package template_method_pattern

import "fmt"

type RuneDisplay struct {
	ch rune
}

func NewRuneDisplay(ch rune) AbstractDisplay {
	return &RuneDisplay{ch: ch}
}

func (r RuneDisplay) Open() {
	fmt.Print("<<")
}

func (r RuneDisplay) Print() {
	fmt.Print(string(r.ch))
}

func (r RuneDisplay) Close() {
	fmt.Println(">>")
}
