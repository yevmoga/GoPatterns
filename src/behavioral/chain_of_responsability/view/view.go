package view

import "behavioral/chain_of_responsability/contract"

type view string

func (p view) ToString() string {
	return string(p)
}

func (p view) toHtml() view {
	return "<p>" + p + "</p>"
}

func Do(data *contract.Data) (contract.Stringer, error) {
	v := view(data.Text)

	return v.toHtml(), nil
}

