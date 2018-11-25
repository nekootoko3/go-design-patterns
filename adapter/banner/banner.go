package banner

import "fmt"

type Banner struct {
	str string
}

func New(str string) *Banner {
	return &Banner{str: str}
}

func (b *Banner) ShowWithParen() {
	fmt.Printf("(%s)\n", b.str)
}

func (b *Banner) ShowWithAster() {
	fmt.Printf("*%s*\n", b.str)
}
