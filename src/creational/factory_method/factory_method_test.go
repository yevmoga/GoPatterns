package main

import (
	"testing"
)

func TestFactoryMethod(t *testing.T) {

	var creatorA Creator = ConcreteCreatorA{}
	var creatorB Creator = ConcreteCreatorB{}

	if creatorA.FactoryMethod().Do() != "Product A do smth" {
		t.Fatal("test failed")
	}

	if creatorB.FactoryMethod().Do() != "Product B do smth" {
		t.Fatal("test failed")
	}

}
