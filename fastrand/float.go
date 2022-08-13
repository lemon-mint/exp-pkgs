package fastrand

// Float64 returns a random float64 in [0, 1).
func (rng *RNG) Float64() float64 {
retry:
	f := float64(rng.Int63()) / (1 << 63)
	if f == 1 {
		goto retry
	}
	return f
}

// Float32 returns a random float32 in [0, 1).
func (rng *RNG) Float32() float32 {
retry:
	f := float32(rng.Float64())
	if f == 1 {
		goto retry
	}
	return f
}

func Float64() float64 {
	r := AcquireRNG()
	v := r.Float64()
	ReleaseRNG(r)
	return v
}

func Float32() float32 {
	r := AcquireRNG()
	v := r.Float32()
	ReleaseRNG(r)
	return v
}
