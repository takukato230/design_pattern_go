package template_method_pattern

type DisplayExecutor struct {
	d AbstractDisplay
}

func NewDisplayExecutor(d AbstractDisplay) DisplayExecutor {
	return DisplayExecutor{d: d}
}

func (e DisplayExecutor) Display() {
	e.d.Open()
	for i := 0; i < 5; i++ {
		e.d.Print()
	}
	e.d.Close()
}
