package visitor

import "fmt"

type Visitor interface {
	visitSquare(Square)
	visitCircle(Circle)
}

type XMLExportVisitor struct {}

func (*XMLExportVisitor) visitSquare(s Square) {
	fmt.Println("export square to xml")
}

func (*XMLExportVisitor) visitCircle(c Circle) {
	fmt.Println("export circle to xml")
}
