package decorator

import (
	"testing"
)

func TestDecorator(t *testing.T) {

	testcases := []struct {
		name           string
		provider       Provider
		expectedResult string
	}{
		{"facebook test case", &Facebook{"John"}, "Hello, John"},
		{"sms test case", &Sms{"38093083623"}, "Hello, user with number 38093083623"},
	}

	for _, testcase := range testcases {
		result := testcase.provider.Send("Hello, ")
		if result != testcase.expectedResult {
			t.Fatalf("%s is failed with result %s. Expected result is %s", testcase.name, result, testcase.expectedResult)
		}
	}
}
