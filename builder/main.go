package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/builder/framework"
	"github.com/nekootoko3/go-design-patterns/builder/html_builder"
	"github.com/nekootoko3/go-design-patterns/builder/text_builder"
)

func main() {
	tb := text_builder.NewTextBuilder()
	tbd := framework.NewDirector(tb)
	tbd.Construct()
	fmt.Println(tb.GetResult())

	hb := html_builder.NewHtmlBuilder()
	hbd := framework.NewDirector(hb)
	hbd.Construct()
	hb.Print()
}
