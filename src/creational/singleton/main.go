package main

import "fmt"

func main() {

	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())
	fmt.Println(GetInstance())

}

type connection struct {
	Address string
	Port    string
}

var connector *connection

func GetInstance() *connection {
	if connector == nil {
		fmt.Println("create connection")
		connector = &connection{Address: "127.0.0.1", Port: "8080"}
	}
	fmt.Println("get connection")
	return connector
}
