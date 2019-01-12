package main

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

type WindowsCheckbox struct {
	Name      string
	IsClicked bool
}

func (checkbox *WindowsCheckbox) Click() {
	checkbox.IsClicked = !checkbox.IsClicked
}

func (checkbox *WindowsCheckbox) IsClick() bool {
	return checkbox.IsClicked
}
