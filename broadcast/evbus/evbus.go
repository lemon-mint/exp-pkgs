package evbus

import (
	"strconv"
	"sync"
	"sync/atomic"

	"v8.run/go/exp/broadcast"
)

var evbus_map sync.Map

func Publish[T any](topic string, v T) {
	t, ok := evbus_map.Load(topic)
	if !ok || t == nil {
		t, _ = evbus_map.LoadOrStore(topic, broadcast.NewTopic[any]())
	}
	t.(*broadcast.Topic[any]).Broadcast(v)
}

type Subscription struct {
	topic string
	s     *broadcast.Subscription[any]
	stop  atomic.Uint32
}

func Subscribe[T any](topic string, fn func(v T) (Unsubscribe bool)) *Subscription {
	t, ok := evbus_map.Load(topic)
	if !ok || t == nil {
		t, _ = evbus_map.LoadOrStore(topic, broadcast.NewTopic[any]())
	}
	var ss Subscription
	ss.topic = topic

	ss.s = t.(*broadcast.Topic[any]).Subscribe(func(v any) {
		if ss.stop.Load() == 1 {
			return
		}

		if val, ok := v.(T); ok {
			unsub := fn(val)
			if unsub {
				ss.Unsubscribe()
			}
		}
	})

	return &ss
}

func (ss *Subscription) Unsubscribe() {
	if ss.stop.CompareAndSwap(0, 1) {
		go ss.s.Unsubscribe()
	}
}

func (ss *Subscription) String() string {
	return "<Subscription " + strconv.Quote(ss.topic) + " >"
}
