package decorator

type Sms struct {
	Number string
}

func (s *Sms) Send(what string) string {
	return what + "user with number " + s.Number
}
