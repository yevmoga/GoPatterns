package visitor

type Shape interface {
	Move(int, int)
	Draw()
	Accept(Visitor) string
}

type Circle struct {
	x, y, r int
}

func (c *Circle) Move(x int, y int) {
	c.x = x
	c.y = y
}

func (*Circle) Draw() {
	panic("implement me")
}

func (c *Circle) Accept(v Visitor) string {
	return v.visitCircle(*c)
}

type Square struct {
	x, y, a int
}

func (*Square) Move(int, int) {
	panic("implement me")
}

func (*Square) Draw() {
	panic("implement me")
}

func (s *Square) Accept(v Visitor) string {
	return v.visitSquare(*s)
}
