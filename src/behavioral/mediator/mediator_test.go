package mediator

import (
	"testing"
)

func TestMediator(t *testing.T) {

	mediator := ConcreteMediator{}

	firstComponent := &Component{"1", &mediator}
	secondComponent := &Component{"2", &mediator}
	thirdComponent := &Component{"3", &mediator}

	mediator.Add(firstComponent)
	mediator.Add(secondComponent)
	mediator.Add(thirdComponent)

	thirdComponent.send("hello world")

}
