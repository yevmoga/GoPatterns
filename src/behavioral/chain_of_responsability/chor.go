package chain_of_responsability

import (
	"behavioral/chain_of_responsability/auth"
	"behavioral/chain_of_responsability/contract"
	"behavioral/chain_of_responsability/process"
	"behavioral/chain_of_responsability/validation"
	"behavioral/chain_of_responsability/view"
)

func Do() []func(*contract.Data) (contract.Stringer, error) {

	return []func(*contract.Data) (contract.Stringer, error){
		auth.Do,
		validation.Do,
		process.Do,
		view.Do,
	}
}
