package command

import (
	"fmt"
	"testing"
)

func TestCommand(t *testing.T) {

	type User struct {
		Waiter Orderer
	}

	user := User{Waiter: &Waiter{}}

	user.Waiter.SetOrder("Olivie")
	user.Waiter.SetOrder("Seledka pod shuboy")
	user.Waiter.SetOrder("Shampusik")
	user.Waiter.SetOrder("Napoleon")
	fmt.Println(user.Waiter.Execute())
	if user.Waiter.Execute() != "executed dish Olivie; executed dish Seledka pod shuboy; executed dish Shampusik; executed dish Napoleon; " {
		t.Fatal("test failed")
	}

}
