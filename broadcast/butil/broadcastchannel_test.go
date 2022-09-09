package butil

import (
	"testing"
)

func TestBroadcastChannel(t *testing.T) {
	bc := BroadcastChannel("Hello, World!")
	defer bc.Close()

	var ctr uint32
	bc.OnMessage = func(v interface{}) {
		ctr += v.(uint32)
	}
	for i := 0; i < 100; i++ {
		bc.PostMessage(uint32(100))
	}
	if ctr != 10000 {
		t.Errorf("ctr == %d, want 10000", ctr)
	}
}
