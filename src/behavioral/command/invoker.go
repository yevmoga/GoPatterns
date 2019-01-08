package command

import "fmt"

type User struct {
	Waiter Orderer
}

func Do() {

	user := User{Waiter: &Waiter{}}
	user.Waiter.SetOrder("Olivie")
	user.Waiter.SetOrder("Seledka pod shuboy")
	user.Waiter.SetOrder("Shampusik")
	user.Waiter.SetOrder("Napoleon")
	table := user.Waiter.Execute()
	fmt.Println(table)

}
