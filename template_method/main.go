package main

import "github.com/nekootoko3/go-design-patterns/template_method/template_method"

func main() {
	cd := template_method.NewCharDisplay("neko")
	d1 := template_method.NewAbstractDisplay(cd)
	d1.Display()

	sd1 := template_method.NewStringDisplay("neko")
	d2 := template_method.NewAbstractDisplay(sd1)
	d2.Display()

	sd2 := template_method.NewStringDisplay("ネコ")
	d3 := template_method.NewAbstractDisplay(sd2)
	d3.Display()
}
