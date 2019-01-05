package chain_of_responsability

import (
	"behavioral/chain_of_responsability/auth"
	"behavioral/chain_of_responsability/contract"
	"behavioral/chain_of_responsability/process"
	"behavioral/chain_of_responsability/validation"
	"behavioral/chain_of_responsability/view"
)

func Do() string {
	d := contract.Data{true, "some text"}
	arr := []func(*contract.Data) (contract.Stringer, error){
		auth.Do,
		validation.Do,
		process.Do,
		view.Do,
	}
	result := ""

	for _, fun := range arr {
		funResult, err := fun(&d)
		if err != nil {
			return err.Error()
		}
		result = funResult.ToString()
	}

	return result
}
