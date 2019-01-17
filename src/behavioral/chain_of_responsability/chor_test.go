package chain_of_responsability

import (
	"behavioral/chain_of_responsability/contract"
	"testing"
)

func TestChor(t *testing.T) {

	d := contract.Data{true, "some text"}
	result := ""

	chor := Do()

	for _, fun := range chor {
		funResult, err := fun(&d)
		if err != nil {
			t.Fatalf("some error: %s", err.Error())
		}

		result = funResult.ToString()
	}

	if result != "<p>some text is grate!</p>" {
		t.Fatal("test failed")
	}
}
