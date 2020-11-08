package main

import "design_pattern_go/template_method_pattern"

func main() {
	dRune := template_method_pattern.NewDisplayExecutor(template_method_pattern.NewRuneDisplay('H'))
	dString := template_method_pattern.NewDisplayExecutor(template_method_pattern.NewStringDisplay("hello world"))
	dRune.Display()
	dString.Display()
}
