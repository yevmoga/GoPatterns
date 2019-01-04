package facade

// this is facade
type Man struct {
	c Child
	h House
	t Tree
}

func (man *Man) Todo() string {
	return man.c.Born() + ", " + man.t.Grow() + ", " + man.h.Build()
}

func Do() string {
	man := Man{Child{}, House{}, Tree{}}
	return man.Todo()
}
