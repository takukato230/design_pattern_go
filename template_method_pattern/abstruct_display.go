package template_method_pattern

type AbstractDisplay interface {
	Open()
	Print()
	Close()
}
