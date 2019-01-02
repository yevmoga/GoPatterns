package bridge

type formatter interface {
	ToFormat() string
}

type Xml struct {
	Open  string
	Close string
}

func (xml *Xml) ToFormat() string {
	xml.Open = "<tag>"
	xml.Close = "</tag>"
	return xml.Open + "example" + xml.Close
}

type Json struct {
	Open  string
	Close string
}

func (json *Json) ToFormat() string {
	json.Open = "{"
	json.Close = "}"
	return json.Open + "example: true" + json.Close
}
