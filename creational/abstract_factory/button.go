package main

type Button interface {
	OnClick() string
	OnDoubleClick() string
}

type WindowsButton struct {
	Name      string
	PositionX uint
	PositionY uint
}

func (WindowsButton) OnClick() string {
	return "You click windows button!"
}

func (WindowsButton) OnDoubleClick() string {
	return "You double click windows button!"
}

type LinuxButton struct {
	Name      string
	PositionX uint
	PositionY uint
}

func (LinuxButton) OnClick() string {
	return "You click linux button!"
}

func (LinuxButton) OnDoubleClick() string {
	return "You double click linux button!"
}

type MacButton struct {
	Name      string
	PositionX uint
	PositionY uint
}

func (MacButton) OnClick() string {
	return "You click mac button!"
}

func (MacButton) OnDoubleClick() string {
	return "You double click mac button!"
}
