package state

type Stater interface {
	Do() string
}

type Registration struct{}

func (Registration) Do() string { return "user registration" }

type Authentication struct{}

func (Authentication) Do() string { return "user authentication" }

type Authorization struct{}

func (Authorization) Do() string { return "user authorization" }

type Logout struct{}

func (Logout) Do() string { return "user logout" }
