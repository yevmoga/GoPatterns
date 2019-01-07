package template_method

type Xml struct {
	Default
}

func (s *Xml) Open() string {
	return "<text>"
}

func (s *Xml) Close() string {
	return "</text>"
}
