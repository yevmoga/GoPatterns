package main

import (
	"behavioral/memento"
	"fmt"
)

func main() {
	originator := memento.Originator{}
	caretaker := memento.Caretaker{}

	originator.SetState("On")
	caretaker.Memento = originator.Save()
	fmt.Println(originator)

	originator.SetState("Off")
	fmt.Println(originator)

	originator.Rollback(caretaker.Memento)

	fmt.Println(originator)

}
