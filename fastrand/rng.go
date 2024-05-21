package fastrand

import (
	"crypto/rand"
	"encoding/binary"
	"sync"

	"gopkg.eu.org/exppkgs/fastrand/alg/splitmix64"
	"gopkg.eu.org/exppkgs/util/noescape"
)

type RNG struct {
	state [2]uint64
}

func refill(r *RNG) {
	var data [16]byte
	_, err := noescape.Read(rand.Reader, data[:])
	if err != nil {
		// If crypto/rand fails, we'll just use the runtime.fastrand implementation.
		binary.LittleEndian.PutUint32(data[0:4], runtime_fastrand())
		binary.LittleEndian.PutUint32(data[4:8], runtime_fastrand())
		binary.LittleEndian.PutUint32(data[8:12], runtime_fastrand())
		binary.LittleEndian.PutUint32(data[12:16], runtime_fastrand())
	}

	r.state[0] = binary.LittleEndian.Uint64(data[0:8])
	r.state[1] = binary.LittleEndian.Uint64(data[8:16])

	// Use Splitmix64 to initialize the state.
	r.state[0] += 1757750930446974760
	r.state[1] += 7151402297004559274
	r.state[0] = splitmix64.Splitmix64(&r.state[0])
	r.state[1] = splitmix64.Splitmix64(&r.state[1])
}

func newRNG() *RNG {
	r := new(RNG)
	refill(r)
	return r
}

var rngPool sync.Pool = sync.Pool{
	New: func() interface{} {
		return newRNG()
	},
}

func AcquireRNG() *RNG {
	return rngPool.Get().(*RNG)
}

func ReleaseRNG(r *RNG) {
	rngPool.Put(r)
}

func WithSeed(seed uint64) *RNG {
	r := AcquireRNG()
	r.SetSeed(seed)
	return r
}

func (rng *RNG) SetSeed(seed uint64) {
	rng.state[0] = splitmix64.Splitmix64(&seed)
	rng.state[1] = splitmix64.Splitmix64(&seed)
}

// Release Put the RNG back into the pool.
// After calling this, the RNG is invalid and should not be used.
func (rng *RNG) Release() {
	ReleaseRNG(rng)
}

// Refill Initialize the RNG with a new seed.
func (rng *RNG) Refill() {
	refill(rng)
}

type FastRandReader struct {
	RNG *RNG
}

func (r *FastRandReader) Read(p []byte) (n int, err error) {
	if r.RNG == nil {
		r.RNG = AcquireRNG()
	}
	n = len(p)

	for len(p) >= 8 {
		binary.LittleEndian.PutUint64(p, r.RNG.Uint64())
		p = p[8:]
	}

	if len(p) > 0 {
		v := r.RNG.Uint64()
		for i := 0; i < len(p); i++ {
			v >>= 8
			p[i] = byte(v)
		}
	}

	return
}
