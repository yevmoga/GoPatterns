package observer

import (
	"testing"
)

func TestObserver(t *testing.T) {

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

	if publisher.NotifySubscribers() != "sending wolf email to subscriber Nouf-Nouf" {
		t.Fatal("test failed")
	}

}
