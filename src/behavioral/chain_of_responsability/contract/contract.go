package contract

type Stringer interface {
	ToString() string
}

type Data struct {
	IsAuth bool
	Text   string
}
