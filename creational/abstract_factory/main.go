package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	fmt.Println("Abstract Factory")

	var userGui GuiFactory
	var userButton Button
	var userCheckbox Checkbox

	switch runtime.GOOS {
	case "linux":
		userGui = LinuxFactory{}
	case "windows":
		userGui = WindowsFactory{}
	case "mac":
		// not sure it's really work :)
		userGui = MacFactory{}
	default:
		log.Printf("cannot found your os %s", runtime.GOOS)
		return
	}

	userButton = userGui.CreateButton()
	userCheckbox = userGui.CreateCheckbox()

	fmt.Println("on button click: ", userButton.OnClick())
	fmt.Println("on button double click: ", userButton.OnDoubleClick())

	fmt.Println("is checkbox clicked(on creating): ", userCheckbox.IsClick())
	fmt.Println("checkbox clicking")
	userCheckbox.Click()
	fmt.Println("is checkbox clicked: ", userCheckbox.IsClick())
	fmt.Println("checkbox clicking")
	userCheckbox.Click()
	fmt.Println("is checkbox clicked: ", userCheckbox.IsClick())
	fmt.Println("checkbox clicking")
	userCheckbox.Click()
	fmt.Println("is checkbox clicked: ", userCheckbox.IsClick())

}
