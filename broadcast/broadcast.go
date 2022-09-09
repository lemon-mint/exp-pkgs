package broadcast

import "sync"

type Topic[T any] struct {
	subscriptions []*Subscription[T]
	rwmu          sync.RWMutex
	pool          sync.Pool
}

func NewTopic[T any]() *Topic[T] {
	t := &Topic[T]{}
	t.pool.New = func() any {
		s := &Subscription[T]{
			t:        t,
			callback: nil,
		}
		return s
	}

	return t
}

type Subscription[T any] struct {
	t        *Topic[T]
	callback func(T)
}

func (t *Topic[T]) Subscribe(callback func(v T)) *Subscription[T] {
	s := t.pool.Get().(*Subscription[T])
	s.callback = callback
	s.t = t

	t.rwmu.Lock()
	t.subscriptions = append(t.subscriptions, s)
	t.rwmu.Unlock()

	return s
}

func (t *Topic[T]) Unsubscribe(s *Subscription[T]) {
	t.rwmu.Lock()
	for i := range t.subscriptions {
		if t.subscriptions[i] == s {
			t.subscriptions[i] = t.subscriptions[len(t.subscriptions)-1]
			t.subscriptions[len(t.subscriptions)-1] = nil
			t.subscriptions = t.subscriptions[:len(t.subscriptions)-1]
			break
		}
	}
	t.rwmu.Unlock()

	s.callback = nil
	t.pool.Put(s)

	return
}

func (t *Topic[T]) Subscribers() uint64 {
	t.rwmu.RLock()
	s := uint64(len(t.subscriptions))
	t.rwmu.RUnlock()
	return s
}

func (s *Subscription[T]) Unsubscribe() {
	s.t.Unsubscribe(s)
}

func (s *Subscription[T]) Topic() *Topic[T] {
	return s.t
}

func (t *Topic[T]) Broadcast(v T) {
	t.rwmu.RLock()
	for i := range t.subscriptions {
		t.subscriptions[i].callback(v)
	}
	t.rwmu.RUnlock()
}
