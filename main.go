package main

import (
	"fmt"
	"structural/decorator"
)

func main() {
	fmt.Println(decorator.Do(&decorator.Sms{"380930000001"}))
}
