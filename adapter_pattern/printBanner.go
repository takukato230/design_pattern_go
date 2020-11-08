package adapter_pattern

type PrintBanner struct {
	banner Banner
}

func NewPrint(str string) Print {
	return &PrintBanner{banner: NewBanner(str)}
}

func (p PrintBanner) PrintWeak() {
	p.banner.showWithParen()
}

func (p PrintBanner) PrintStrong() {
	p.banner.showWithAster()
}
