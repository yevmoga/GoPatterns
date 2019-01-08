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

func (p *Publisher) NotifySubscribers() {
	for _, single := range p.subscribers {
		single.Update()
	}
}

func Do() {

	subscribers := []subscriber{
		Subscriber{"Nif-Nif"},
		Subscriber{"Naf-Naf"},
		Subscriber{"Nouf-Nouf"},
	}

	publisher := Publisher{}
	publisher.subscribers = make(map[string]subscriber)
	for _, s := range subscribers {
		publisher.Subscribe(s)
	}

	publisher.Unsubscribe(subscribers[0])
	publisher.Unsubscribe(subscribers[1])

	publisher.NotifySubscribers()

}
