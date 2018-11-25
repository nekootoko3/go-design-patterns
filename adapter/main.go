package main

import (
	"fmt"

	"github.com/nekootoko3/go-design-patterns/adapter/print"

	dele_pb "github.com/nekootoko3/go-design-patterns/adapter/dele/print_banner"
	inhr_pb "github.com/nekootoko3/go-design-patterns/adapter/inhr/print_banner"
)

func main() {
	var p print.Print

	fmt.Println("inheritance? ---------")
	p = inhr_pb.New("neko")
	p.PrintWeak()
	p.PrintStrong()

	fmt.Println("delegation? ---------")
	p = dele_pb.New("neko")
	p.PrintWeak()
	p.PrintStrong()
}
