package memento

type Originator struct {
	state string
}

func (o *Originator) Save() *Memento {
	return &Memento{state: o.state}
}

func (o *Originator) Rollback(memento *Memento) {
	o.state = memento.GetState()
}

func (o *Originator) SetState(state string) {
	o.state = state
}
