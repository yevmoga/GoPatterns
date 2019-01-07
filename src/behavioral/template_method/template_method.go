package template_method

type Structure interface {
	Open() string
	Context() string
	Close() string
}

type Default struct{}

func (s *Default) Context() string {
	return "some default text"
}
