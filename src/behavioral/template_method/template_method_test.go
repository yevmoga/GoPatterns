package template_method

import (
	"testing"
)

func TestTemplateMethod(t *testing.T) {
	testCases := []struct{
		format Structure
		result string
	} {
		{&Xml{}, "<text>some default text</text>"},
		{&Json{}, "{some default text}"},
	}

	for _, testCase := range testCases {
		if testCase.format.Open() + testCase.format.Context() + testCase.format.Close() != testCase.result {
			t.Fatal("error test case")
		}
	}
}
