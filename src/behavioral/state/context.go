package state

type Context struct {
	state Stater
}

func (c *Context) Do() string {
	return c.state.Do()
}

func (c *Context) SetState(state Stater) {
	c.state = state
}
