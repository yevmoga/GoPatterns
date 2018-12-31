package adapter

import (
	"fmt"
	"structural/adapter/remote"
)

func Do() {

	var remoteUser = remote.GetUsefulInfo()
	var u CustomerAdapter = &User{}

	fmt.Println(u.ToUser(remoteUser))
	fmt.Println(u.ToRemoteService())
}
