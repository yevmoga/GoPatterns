package main

type Prototyper interface {
	SetName(string)
	GetName() string
	Clone() Prototyper
}

type PrototypeOne struct {
	Name string
}

func (prototype *PrototypeOne) SetName(name string) {
	prototype.Name = name
}
func (prototype PrototypeOne) GetName() string {
	return prototype.Name
}

func (prototype PrototypeOne) Clone() Prototyper {
	return &prototype
}
