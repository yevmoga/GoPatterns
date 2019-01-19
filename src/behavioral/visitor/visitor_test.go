package visitor

import "testing"

func TestVisitor(t *testing.T) {

	xmlExporter := XMLExportVisitor{}
	circle := Circle{}
	square := Square{}

	if circle.Accept(&xmlExporter) != "export circle to xml" {
		t.Fatal("circle test failed")
	}

	if square.Accept(&xmlExporter) != "export square to xml" {
		t.Fatal("square test failed")
	}

}
