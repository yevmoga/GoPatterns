package composite

import (
	"testing"
)

func TestComposite(t *testing.T) {

	composite := composite{
		price: 1,
		composits: []compositer{
			&composite{price: 2},
			&composite{composits: []compositer{
				&composite{price: 3},
				&composite{price: 4},
			},},
		},
	}

	if composite.GetPrice() != 10 {
		t.Fatal("error testing")
	}

}
