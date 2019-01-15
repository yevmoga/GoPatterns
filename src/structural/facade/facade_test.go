package facade

import "testing"

func TestFacade(t *testing.T) {
	man := Man{Child{}, House{}, Tree{}}

	if man.c.Born() != "child born" {
		t.Fatal("not born")
	}

	if man.t.Grow() != "tree grow" {
		t.Fatal("not grow")
	}

	if man.h.Build() != "build house" {
		t.Fatal("not build")
	}
}
