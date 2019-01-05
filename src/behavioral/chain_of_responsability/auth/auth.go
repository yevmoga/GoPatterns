package auth

import (
	"behavioral/chain_of_responsability/contract"
	"errors"
)

type isAuth string
func (ia isAuth) ToString() string {
	return string(ia)
}

func Do(data *contract.Data) (contract.Stringer, error) {
	if !data.IsAuth {
		return isAuth("Error!"), errors.New("user not found")
	}

	return isAuth("Ok!"), nil
}
