package main

type Checkbox interface {
	Click()
	IsClick() bool
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
