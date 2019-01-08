package command

type Orderer interface {
	Execute() string
	SetOrder(string)
}

type Waiter struct {
	Orders   []string
	Receiver Chef
}

func (c *Waiter) SetOrder(s string) {
	c.Orders = append(c.Orders, s)
}

func (c *Waiter) Execute() string {
	result := ""
	for _, order := range c.Orders {
		result += c.Receiver.Execute(order) + "; "
	}
	return result
}
