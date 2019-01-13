package proxy

import (
	"testing"
)

func TestProxy(t *testing.T) {
	input := "call service"
	proxy := CreateProxy(&Service{})
	result := proxy.Get(input)
	if result != input + " with some proxy" {
		t.Fatal("test failed")
	}
}
