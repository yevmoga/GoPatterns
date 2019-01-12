package main

import (
	"log"
	"runtime"
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	type action struct {
		buttonResults   []string
		checkboxResults []bool
	}

	testCases := []struct {
		os     string
		action action
	}{
		{"linux", action{[]string{"You click linux button!", "You double click linux button!"}, []bool{false, true, false, true,}}},
		{"windows", action{[]string{"You click windows button!", "You double click windows button!"}, []bool{false, true, false, true,}}},
		{"mac", action{[]string{"You click mac button!", "You double click mac button!"}, []bool{false, true, false, true,}}},
	}

	var userGui GuiFactory
	var userButton Button
	var userCheckbox Checkbox

	for _, testcase := range testCases {

		switch testcase.os {
		case "linux":
			userGui = LinuxFactory{}
		case "windows":
			userGui = WindowsFactory{}
		case "mac":
			userGui = MacFactory{}
		default:
			log.Printf("cannot found your os %s", runtime.GOOS)
			return
		}

		userButton = userGui.CreateButton()
		userCheckbox = userGui.CreateCheckbox()

		if userButton.OnClick() != testcase.action.buttonResults[0] {
			t.Fatal("test failed")
		}

		if userButton.OnDoubleClick() != testcase.action.buttonResults[1] {
			t.Fatal("test failed")
		}

		if userCheckbox.IsClick() != testcase.action.checkboxResults[0] {
			t.Fatal("test failed")
		}

		if userCheckbox.Click(); userCheckbox.IsClick() != testcase.action.checkboxResults[1] {
			t.Fatal("test failed")
		}

		if userCheckbox.Click(); userCheckbox.IsClick() != testcase.action.checkboxResults[2] {
			t.Fatal("test failed")
		}

		if userCheckbox.Click(); userCheckbox.IsClick() != testcase.action.checkboxResults[3] {
			t.Fatal("test failed")
		}
	}

}
