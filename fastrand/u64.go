package fastrand

// This file contains the implementation of the xorshift128+ algorithm.
// The algorithm is described in: https://vigna.di.unimi.it/ftp/papers/xorshiftplus.pdf
func xorshift128plus(state *[2]uint64) uint64 {
	var x, y, result uint64
	x, y = state[1], state[0]
	result = x + y
	state[0] = x
	y ^= y << 23
	state[1] = x ^ y ^ (x >> 5) ^ (y >> 18)
	return result
}

func (rng *RNG) Uint64() uint64 {
	return xorshift128plus(&rng.state)
}

func (rng *RNG) Uint64n(n uint64) uint64 {
	return xorshift128plus(&rng.state) % n
}

func (rng *RNG) Int64() int64 {
	return int64(rng.Uint64())
}

func (rng *RNG) Int63() int64 {
	return int64(rng.Uint64() & (1<<63 - 1))
}

func (rng *RNG) Int63n(n int64) int64 {
	return int64(rng.Uint64n(uint64(n)))
}

func Uint64() uint64 {
	r := AcquireRNG()
	v := r.Uint64()
	ReleaseRNG(r)
	return v
}

func Uint64n(n uint64) uint64 {
	r := AcquireRNG()
	v := r.Uint64n(n)
	ReleaseRNG(r)
	return v
}

func Int64() int64 {
	r := AcquireRNG()
	v := r.Int64()
	ReleaseRNG(r)
	return v
}

func Int63() int64 {
	r := AcquireRNG()
	v := r.Int63()
	ReleaseRNG(r)
	return v
}

func Int63n(n int64) int64 {
	r := AcquireRNG()
	v := r.Int63n(n)
	ReleaseRNG(r)
	return v
}
