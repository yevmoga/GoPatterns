package main

type connection struct {
	Address string
	Port    string
}

var connector *connection

func GetInstance() *connection {
	if connector == nil {
		connector = &connection{Address: "127.0.0.1", Port: "8080"}
	}
	return connector
}
