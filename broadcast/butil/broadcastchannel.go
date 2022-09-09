package butil

import (
	"sync"

	"v8.run/go/exp/broadcast"
)

var broadcastChannelMap = sync.Map{}

type BroadcastChannelTopic struct {
	s         *broadcast.Subscription[interface{}]
	name      string
	OnMessage func(v interface{})
}

func BroadcastChannel(name string) *BroadcastChannelTopic {
	v, ok := broadcastChannelMap.Load(name)
	if v != ok {
		v, _ = broadcastChannelMap.LoadOrStore(name, broadcast.NewTopic[interface{}]())
	}

	bt := &BroadcastChannelTopic{
		name: name,
	}

	s := v.((*broadcast.Topic[interface{}])).Subscribe(func(v interface{}) {
		o := bt.OnMessage
		if o != nil {
			o(v)
		}
	})
	bt.s = s

	return bt
}

func (bt BroadcastChannelTopic) Name() string {
	return bt.name
}

func (bt BroadcastChannelTopic) PostMessage(msg interface{}) {
	bt.s.Topic().Broadcast(msg)
	return
}

func (bt BroadcastChannelTopic) Close() {
	bt.s.Unsubscribe()
}
