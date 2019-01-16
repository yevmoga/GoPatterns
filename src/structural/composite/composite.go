package composite

type compositer interface {
	GetPrice() int
}

type composite struct {
	price     int
	composits []compositer
}

func (currentComposite *composite) GetPrice() int {
	if currentComposite.composits != nil {
		for _, composite := range currentComposite.composits {
			currentComposite.price += composite.GetPrice()
		}
	}
	return currentComposite.price
}
