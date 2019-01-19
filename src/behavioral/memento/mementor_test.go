package memento

import (
	"testing"
)

func TestMemento(t *testing.T) {

	originator := Originator{}
	caretaker := Caretaker{}

	originator.SetState("On")
	caretaker.Memento = originator.Save()
	if originator.state != "On" {
		t.Fatal("test failed")
	}

	originator.SetState("Off")
	if originator.state != "Off" {
		t.Fatal("test failed")
	}

	originator.Rollback(caretaker.Memento)
	if originator.state != "On" {
		t.Fatal("test failed")
	}
}
