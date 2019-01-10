package mediator

type Mediator interface {
	Send(message string, currentComponent Componenter)
	Add(component Componenter)
}

type ConcreteMediator struct {
	Components []Componenter
}

func (c *ConcreteMediator) Send(message string, currentComponent Componenter) {
	for _, component := range c.Components {
		if component != currentComponent {
			component.receive(message)
		}
	}
}

func (c *ConcreteMediator) Add(component Componenter) {
	c.Components = append(c.Components, component)
}
