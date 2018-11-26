package underline_pen

import (
	"fmt"
	"strings"

	"github.com/nekootoko3/go-design-patterns/prototype/framework"
)

type underlinePen struct {
	ulChar string
}

func New(str string) *underlinePen {
	return &underlinePen{ulChar: str}
}

func (ulp *underlinePen) Use(str string) {
	underline := strings.Repeat(ulp.ulChar, len(str))
	fmt.Println(str)
	fmt.Println(underline)
}

func (ulp *underlinePen) CreateClone() framework.Producter {
	return &underlinePen{ulChar: ulp.ulChar}
}
