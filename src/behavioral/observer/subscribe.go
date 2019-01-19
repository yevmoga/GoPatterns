package observer

import "fmt"

type subscriber interface {
	GetName() string
	Update() string
}

type Subscriber struct {
	Name string
}

func (s Subscriber) Update() string {
	return fmt.Sprintf("sending wolf email to subscriber %s", s.Name)
}

func (s Subscriber) GetName() string {
	return s.Name
}
