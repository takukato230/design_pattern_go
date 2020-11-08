package main

import "design_pattern_go/adapter_pattern"

func main() {
	p := adapter_pattern.NewPrint("Hello")
	p.PrintStrong()
	p.PrintWeak()
}
