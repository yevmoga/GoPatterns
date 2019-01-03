package decorator

type Facebook struct {
	FriendName string
}

func (f *Facebook) Send(what string) string {
	return what + f.FriendName
}
