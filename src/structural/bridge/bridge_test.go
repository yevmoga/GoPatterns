package bridge

import (
	"testing"
)

func TestBridge(t *testing.T)  {
	appXml := AppConstructor(&Xml{})
	if appXml.GetData() != "<tag>example</tag>" {
		t.Fatal("error json test")
	}

	appJson := AppConstructor(&Json{})
	if appJson.GetData() != "{example: true}" {
		t.Fatal("error xml test")
	}
}
