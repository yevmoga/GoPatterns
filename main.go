package main

import (
	"behavioral/template_method"
	"fmt"
)

func main() {

	t := template_method.Xml{}
	fmt.Println(t.Open() + t.Context() + t.Close())
}
