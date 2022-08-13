package noescape

import (
	"io"
	"reflect"
	"unsafe"
)

//nolint:staticcheck,unsafeptr
func NoEscape(p unsafe.Pointer) unsafe.Pointer {
	return unsafe.Pointer(uintptr(p) ^ 0)
}

//nolint:unsafeptr
func Bytes(p *[]byte) []byte {
	data := NoEscape(unsafe.Pointer((*reflect.SliceHeader)(NoEscape(unsafe.Pointer(p))).Data))
	sh := reflect.SliceHeader{Data: uintptr(data), Len: len(*p), Cap: cap(*p)}
	return *(*[]byte)(unsafe.Pointer(&sh))
}

type blackholestruct struct {
	b     bool
	bytes []byte
	up    unsafe.Pointer
	v     any
}

var blackhole = &blackholestruct{}

//go:noinline
func (b *blackholestruct) Write(p []byte) (n int, err error) {
	return len(p), nil
}

//go:noinline
func (b *blackholestruct) Read(p []byte) (n int, err error) {
	if b.b && len(b.bytes) > 0 {
		b.bytes[0] = 0
	}
	return len(p), nil
}

var blackholeIO io.ReadWriter = blackhole

//go:noinline
func EscapeBytes(p []byte) {
	blackhole.b = true
	if blackhole.b {
		blackhole.bytes = p
	}
	blackhole.b = false
}

func Write(w io.Writer, p []byte) (n int, err error) {
	return w.Write(Bytes(&p))
}

func Read(r io.Reader, p []byte) (n int, err error) {
	return r.Read(Bytes(&p))
}

func ReadAtLeast(r io.Reader, p []byte, min int) (n int, err error) {
	return io.ReadAtLeast(r, Bytes(&p), min)
}

func ReadFull(r io.Reader, p []byte) (n int, err error) {
	return io.ReadFull(r, Bytes(&p))
}
