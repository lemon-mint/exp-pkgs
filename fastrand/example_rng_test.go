package fastrand_test

import (
	"fmt"

	"v8.run/go/exp/fastrand"
)

func Example_rNG() {
	rng := fastrand.AcquireRNG()
	defer fastrand.ReleaseRNG(rng)

	rng.SetSeed(42)

	fmt.Printf("Number between 0 and 100: %d\n", rng.Int63n(100))
	fmt.Printf("64-bit unsigned integer: %d\n", rng.Uint64())
}
