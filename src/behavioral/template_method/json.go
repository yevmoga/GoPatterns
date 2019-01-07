package template_method

type Json struct {
	Default
}

func (s *Json) Open() string {
	return "{"
}

func (s *Json) Close() string {
	return "}"
}
