package bridge

type Xml struct {
	Open  string
	Close string
}

func (xml *Xml) ToFormat() string {
	xml.Open = "<tag>"
	xml.Close = "</tag>"
	return xml.Open + "example" + xml.Close
}
