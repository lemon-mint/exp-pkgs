package noescape

import (
	"testing"
)

func BenchmarkNoEscapeBytes(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		var b [8]byte
		t := b[:]
		var v []byte
		for p.Next() {
			v = Bytes(&t)
		}
		EscapeBytes(v)
	})
}

func BenchmarkNoEscapeWrite(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			var b [256]byte
			t := b[:]
			Write(blackholeIO, t)
		}
	})
}

func BenchmarkNoEscapeRead(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			var b [256]byte
			t := b[:]
			Read(blackholeIO, t)
		}
	})
}

func BenchmarkWrite(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			var b [256]byte
			t := b[:]
			blackholeIO.Write(t)
		}
	})
}

func BenchmarkRead(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			var b [256]byte
			t := b[:]
			blackholeIO.Read(t)
		}
	})
}
