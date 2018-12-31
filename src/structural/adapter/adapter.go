package adapter

import (
	"fmt"
	"strings"
	"structural/adapter/remote"
	"time"
)

type CustomerAdapter interface {
	ToRemoteService() remote.CustomerInfo
	ToUser(remote.CustomerInfo) string
}

type User struct {
	FullName string
	Age      uint16
}

func (user *User) ToRemoteService() remote.CustomerInfo {
	names := strings.Split(user.FullName, " ")
	return remote.CustomerInfo{
		FirstName:  names[0],
		SecondName: names[1],
		BirthYear:  uint16(time.Now().AddDate(int(-user.Age), 0, 0).Year()),
	}
}

func (user *User) ToUser(userInfo remote.CustomerInfo) string {
	user.FullName = userInfo.FirstName + " " + userInfo.SecondName
	user.Age = uint16(time.Now().AddDate(int(-userInfo.BirthYear), 0, 0).Year())

	return fmt.Sprintf("user full name is %s, and age is %d", user.FullName, user.Age)
}
