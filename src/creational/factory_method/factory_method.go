package main

type Productor interface {
	Do() string
}

type Creator interface {
	FactoryMethod() Productor
}

type ConcreteCreatorA struct {
	Name string
}

type ConcreteCreatorB struct {
	Name string
}

type ConcreteProductA struct {
	Name string
}

type ConcreteProductB struct {
	Name string
}

func (product ConcreteProductA) Do() string {
	return product.Name + " do smth"
}

func (product ConcreteProductB) Do() string {
	return product.Name + " do smth"
}

func (ConcreteCreatorA) FactoryMethod() Productor {
	return ConcreteProductA{"Product A"}
}

func (ConcreteCreatorB) FactoryMethod() Productor {
	return ConcreteProductB{"Product B"}
}
