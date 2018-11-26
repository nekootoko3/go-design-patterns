package message_box

import (
	"fmt"
	"strings"

	"github.com/nekootoko3/go-design-patterns/prototype/framework"
)

type messageBox struct {
	decoStr string
}

func New(str string) *messageBox {
	return &messageBox{decoStr: str}
}

func (mb *messageBox) Use(str string) {
	length := len(str)
	decoLine := strings.Repeat(mb.decoStr, length+2)
	fmt.Println(decoLine)
	fmt.Printf("%s%s%s\n", mb.decoStr, str, mb.decoStr)
	fmt.Println(decoLine)
}

func (mb *messageBox) CreateClone() framework.Producter {
	return &messageBox{decoStr: mb.decoStr}
}
