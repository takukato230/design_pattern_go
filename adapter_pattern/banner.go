package adapter_pattern

import "fmt"

type Banner struct {
	Str string
}

func NewBanner(str string) Banner {
	return Banner{Str: str}
}

func (b Banner) showWithParen() {
	fmt.Printf("(%s)\n", b.Str)
}

func (b Banner) showWithAster() {
	fmt.Printf("*%s*\n", b.Str)
}
