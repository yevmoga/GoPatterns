package observer

import "fmt"

type subscriber interface {
	GetName() string
	Update()
}

type Subscriber struct {
	Name string
}

func (s Subscriber) Update() {
	fmt.Println("sending wolf email")
}

func (s Subscriber) GetName() string {
	return s.Name
}
