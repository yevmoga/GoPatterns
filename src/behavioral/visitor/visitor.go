package visitor

type Visitor interface {
	visitSquare(Square) string
	visitCircle(Circle) string
}

type XMLExportVisitor struct {}

func (*XMLExportVisitor) visitSquare(s Square) string{
	return "export square to xml"
}

func (*XMLExportVisitor) visitCircle(c Circle) string{
	return "export circle to xml"
}
