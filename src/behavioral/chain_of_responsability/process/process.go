package process

import (
	"behavioral/chain_of_responsability/contract"
)

type process string

func (p process) ToString() string {
	return string(p)
}

func Do(data *contract.Data) (contract.Stringer, error) {
	data.Text = data.Text + " is grate!"
	return process(""), nil
}

