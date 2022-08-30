package runtime2

import "testing"

func TestSetGet(t *testing.T) {
	var mls = NewMLocal[int]()
	for i := 0; i < 100000; i++ {
		mls.Set(i)
		if v := mls.Get(); v != i {
			panic("i != v")
		}
	}
}
