package composite

import "fmt"

func Do() {

	composite := composite{
		price:     1,
		composits: []compositer{
			&composite{price:2},
			&composite{composits: []compositer{
				&composite{price:3},
				&composite{price:4},
			},},
		},
	}

	fmt.Println(composite.GetPrice())

}
