package main

type LinuxFactory struct{}

func (LinuxFactory) CreateButton() Button {
	return LinuxButton{
		Name:      "Hello, Linux!",
		PositionX: 150,
		PositionY: 75,
	}
}
func (LinuxFactory) CreateCheckbox() Checkbox {
	return &LinuxCheckbox{
		Name:      "Linux CheckBox",
		IsClicked: false,
	}
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

type LinuxCheckbox struct {
	Name      string
	IsClicked bool
}

func (checkbox *LinuxCheckbox) Click() {
	checkbox.IsClicked = !checkbox.IsClicked
}

func (checkbox *LinuxCheckbox) IsClick() bool {
	return checkbox.IsClicked
}
