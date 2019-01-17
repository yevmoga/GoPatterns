package adapter

import (
	"structural/adapter/remote"
	"testing"
)

func TestAdapter(t *testing.T) {

	var remoteUser = remote.GetUsefulInfo()
	var u CustomerAdapter = &User{}

	if u.ToUser(remoteUser) != "user full name is Michael Jordan, and age is 56" {
		t.Fatal("error user info")
	}

	customer := remote.CustomerInfo{
		FirstName:  "Michael",
		SecondName: "Jordan",
		BirthYear:  1963,
	}

	if u.ToRemoteService() != customer {
		t.Fatal("error user info")
	}
}
