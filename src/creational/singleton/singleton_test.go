package main

import (
	"testing"
)

func TestSingleton(t *testing.T) {
	for i := 0; i < 3; i++ {
		connection := GetInstance()
		if connection.Address != "127.0.0.1" || connection.Port != "8080" {
			t.Fatal("test failed")
		}
	}
}
