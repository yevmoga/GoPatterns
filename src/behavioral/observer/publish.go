package observer

type Publisher struct {
	subscribers map[string]subscriber
	mainState   bool
}

func (p *Publisher) Subscribe(subscriber subscriber) {
	p.subscribers[subscriber.GetName()] = subscriber
}

func (p *Publisher) Unsubscribe(subscriber subscriber) {
	delete(p.subscribers, subscriber.GetName())
}

func (p *Publisher) NotifySubscribers() string {
	var whatHappen string
	for _, single := range p.subscribers {
		whatHappen += single.Update()
	}

	return whatHappen
}
