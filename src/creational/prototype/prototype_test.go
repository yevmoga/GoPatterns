package main

import (
	"testing"
)

func TestPrototype(t *testing.T) {

	p1 := PrototypeOne{}
	p1.SetName("One")
	p2 := p1.Clone()
	p2.SetName("Two")

	if p1.GetName() != "One" {
		t.Fatal("test failed")
	}

	if p2.GetName() != "Two" {
		t.Fatal("test failed")
	}
}
