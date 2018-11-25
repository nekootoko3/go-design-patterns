package print_banner

import "github.com/nekootoko3/go-design-patterns/adapter/banner"

type printBanner struct {
	banner *banner.Banner
}

func New(str string) *printBanner {
	banner := banner.New(str)
	return &printBanner{banner: banner}
}

func (pb *printBanner) PrintWeak() {
	pb.banner.ShowWithParen()
}

func (pb *printBanner) PrintStrong() {
	pb.banner.ShowWithAster()
}
