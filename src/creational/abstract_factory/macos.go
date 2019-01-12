package main

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

type MacCheckbox struct {
	Name      string
	IsClicked bool
}

func (checkbox *MacCheckbox) Click() {
	checkbox.IsClicked = !checkbox.IsClicked
}

func (checkbox *MacCheckbox) IsClick() bool {
	return checkbox.IsClicked
}
