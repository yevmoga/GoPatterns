package main

import (
	"behavioral/visitor"
)

func main() {
	xmlExporter := visitor.XMLExportVisitor{}
	circle := visitor.Circle{}
	circle.Accept(&xmlExporter)
}
