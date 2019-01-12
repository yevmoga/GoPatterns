package main

type GuiFactory interface {
	CreateButton() Button
	CreateCheckbox() Checkbox
}
