package mediator

import "fmt"

type Componenter interface {
	receive(string)
	send(string)
}

type Component struct {
	Name     string
	Mediator Mediator
}

func (c *Component) receive(s string) {
	fmt.Printf("%s got the message \"%s\"\n", c.Name, s)
}

func (c *Component) send(message string) {
	c.Mediator.Send(fmt.Sprintf("\"%s\" message from component %s", message, c.Name), c)
}
