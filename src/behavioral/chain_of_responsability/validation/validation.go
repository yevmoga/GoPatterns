package validation

import (
	"behavioral/chain_of_responsability/contract"
	"errors"
	"strings"
)

type isValidate string

func (iv isValidate) ToString() string {
	return string(iv)
}

func Do(data *contract.Data) (contract.Stringer, error) {
	if len(data.Text) <= 0 {
		return isValidate("Error!"), errors.New("your text is empty")
	}

	if !strings.ContainsRune(data.Text, ' ') {
		return isValidate("Error!"), errors.New("should be more than 2 words")
	}

	return isValidate("Ok!"), nil
}
