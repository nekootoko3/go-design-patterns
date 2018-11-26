package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/prototype/framework"
	"github.com/nekootoko3/go-design-patterns/prototype/message_box"
	"github.com/nekootoko3/go-design-patterns/prototype/underline_pen"
)

func main() {
	mb1 := message_box.New("/")
	mb1.Use("neko")
	fmt.Println("")
	mb2 := message_box.New("~")
	mb2.Use("neko")
	fmt.Println("")
	ulp := underline_pen.New("=")
	ulp.Use("neko")
	fmt.Println("")

	manager := framework.NewManager()
	manager.Register("naname message", mb1)
	manager.Register("nami message", mb2)
	manager.Register("double under", ulp)

	mb1Clone := manager.Clone("naname message")
	mb1Clone.Use("neko")
	fmt.Println("")

	fmt.Printf("orig address: %p\nclone address: %p\n", mb1, mb1Clone)
}
