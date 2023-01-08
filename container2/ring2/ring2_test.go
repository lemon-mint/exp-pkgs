package ring2_test

import (
	"testing"

	"v8.run/go/exp/container2/ring2"
)

func TestRing2(t *testing.T) {
	r := ring2.NewRing2[uint64](3)
	if r.Len() != 0 {
		t.Errorf("expected len 0, got %d", r.Len())
	}

	if r.Cap() != 3 {
		t.Errorf("expected cap 3, got %d", r.Cap())
	}

	if r.Free() != 3 {
		t.Errorf("expected free 3, got %d", r.Free())
	}

	if !r.Write(1) {
		t.Errorf("expected write to succeed")
	}

	if r.Len() != 1 {
		t.Errorf("expected len 1, got %d", r.Len())
	}

	if r.Free() != 2 {
		t.Errorf("expected free 2, got %d", r.Free())
	}

	if !r.Write(2) {
		t.Errorf("expected write to succeed")
	}

	if !r.Write(3) {
		t.Errorf("expected write to succeed")
	}

	if r.Write(4) {
		t.Errorf("expected write to fail")
	}

	data, ok := r.Read()
	if !ok {
		t.Errorf("expected read to succeed")
	}

	if data != 1 {
		t.Errorf("expected data 1, got %d", data)
	}

	if r.Len() != 2 {
		t.Errorf("expected len 2, got %d", r.Len())
	}

	if r.Free() != 1 {
		t.Errorf("expected free 1, got %d", r.Free())
	}

	if !r.Write(4) {
		t.Errorf("expected write to succeed")
	}

	if r.Write(5) {
		t.Errorf("expected write to fail")
	}

	data, ok = r.Read()
	if !ok {
		t.Errorf("expected read to succeed")
	}

	if data != 2 {
		t.Errorf("expected data 2, got %d", data)
	}

	data, ok = r.Read()
	if !ok {
		t.Errorf("expected read to succeed")
	}

	if data != 3 {
		t.Errorf("expected data 3, got %d", data)
	}

	data, ok = r.Read()
	if !ok {
		t.Errorf("expected read to succeed")
	}

	if data != 4 {
		t.Errorf("expected data 4, got %d", data)
	}

	if r.Len() != 0 {
		t.Errorf("expected len 0, got %d", r.Len())
	}

	if r.Free() != 3 {
		t.Errorf("expected free 3, got %d", r.Free())
	}

	_, ok = r.Read()
	if ok {
		t.Errorf("expected read to fail")
	}
}

func TestRing2At(t *testing.T) {
	r := ring2.NewRing2[uint64](3)
	r.Write(1)
	r.Write(2)
	r.Write(3)
	r.Read()
	r.Read()
	r.Write(4)
	r.Write(5)

	if r.At(0) != 3 {
		t.Errorf("expected data 3, got %d", r.At(0))
	}

	if r.At(1) != 4 {
		t.Errorf("expected data 4, got %d", r.At(1))
	}

	if r.At(2) != 5 {
		t.Errorf("expected data 5, got %d", r.At(2))
	}
}

func TestRing2Drop(t *testing.T) {
	r := ring2.NewRing2[uint64](3)
	r.Write(1)
	r.Write(2)
	r.Write(3)
	r.Read()
	r.Read()
	r.Write(4)
	r.Write(5)
	r.Drop(2)

	if r.Len() != 1 {
		t.Errorf("expected len 1, got %d", r.Len())
	}

	if r.Free() != 2 {
		t.Errorf("expected free 2, got %d", r.Free())
	}

	if r.At(0) != 5 {
		t.Errorf("expected data 5, got %d", r.At(0))
	}

	r.Reset()

	if r.Len() != 0 {
		t.Errorf("expected len 0, got %d", r.Len())
	}

	if r.Free() != 3 {
		t.Errorf("expected free 3, got %d", r.Free())
	}

	r.Write(100)
	r.Write(200)
	r.Write(300)
	r.Drop(1)

	if r.Len() != 2 {
		t.Errorf("expected len 2, got %d", r.Len())
	}

	if r.Free() != 1 {
		t.Errorf("expected free 1, got %d", r.Free())
	}

	if r.At(0) != 200 {
		t.Errorf("expected data 200, got %d", r.At(0))
	}

	if r.At(1) != 300 {
		t.Errorf("expected data 300, got %d", r.At(1))
	}
}
