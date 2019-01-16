package bridge

type Json struct {
	Open  string
	Close string
}

func (json *Json) ToFormat() string {
	json.Open = "{"
	json.Close = "}"
	return json.Open + "example: true" + json.Close
}
