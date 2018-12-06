package main

import "fmt"

func main() {
	p1 := PrototypeOne{}
	p1.SetName("One")
	p2 := p1.Clone()
	p2.SetName("Two")
	fmt.Println(p1.GetName())
	fmt.Println(p2.GetName())
}

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
