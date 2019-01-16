package flyweight

import (
	"testing"
)

func TestFlyweight(t *testing.T) {
	flyweightFactory := FlyweightFactory{}
	flyweightFactory.GetFlyweight(100).SetName("Name 1")
	flyweightFactory.GetFlyweight(200).SetName("Name 2")
	flyweightFactory.GetFlyweight(300).SetName("Name 3")

	if flyweightFactory.GetFlyweight(100).GetName() != "Name 1" {
		t.Fatal("error test. Should be Name 1")
	}

	if flyweightFactory.GetFlyweight(200).GetName() != "Name 2" {
		t.Fatal("error test. Should be Name 2")
	}

	if flyweightFactory.GetFlyweight(300).GetName() != "Name 3" {
		t.Fatal("error test. Should be Name 3")
	}

	if flyweightFactory.GetFlyweight(100).GetName() != "Name 1" {
		t.Fatal("error test. Should be Name 1")
	}
}
