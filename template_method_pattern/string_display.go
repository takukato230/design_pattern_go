package template_method_pattern

import (
	"fmt"
	"unicode/utf8"
)

type StringDisplay struct {
	str   string
	width int
}

func NewStringDisplay(str string) AbstractDisplay {
	return &StringDisplay{
		str:   str,
		width: utf8.RuneCountInString(str),
	}
}

func (s StringDisplay) Open() {
	s.printLine()
}

func (s StringDisplay) Print() {
	fmt.Printf("|%s|\n", s.str)
}

func (s StringDisplay) Close() {
	s.printLine()
}

func (s StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
