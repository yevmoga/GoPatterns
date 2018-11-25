package main

type GuiFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}

type WindowsFactory struct{}

func (WindowsFactory) CreateButton() Button {
	return WindowsButton{
		Name:      "Hello, Windows!",
		PositionX: 100,
		PositionY: 50,
	}
}
func (WindowsFactory) CreateCheckbox() Checkbox {
	return &WindowsCheckbox{
		Name:      "Windows CheckBox",
		IsClicked: false,
	}
}

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

type MacFactory struct{}

func (MacFactory) CreateButton() Button {
	return MacButton{
		Name:      "Hello, Mac!",
		PositionX: 120,
		PositionY: 90,
	}
}
func (MacFactory) CreateCheckbox() Checkbox {
	return &MacCheckbox{
		Name:      "Mac CheckBox",
		IsClicked: false,
	}
}
