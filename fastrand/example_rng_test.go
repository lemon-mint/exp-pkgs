package fastrand_test

import (
	"fmt"

	"v8.run/go/exp/fastrand"
)

func Example_rNG() {
	rng := fastrand.AcquireRNG()
	defer fastrand.ReleaseRNG(rng)

	rng.SetSeed(42)

	fmt.Printf("Number between 0 and 10: %d\n", rng.Int63n(10))
	fmt.Printf("Number between 0 and 100: %d\n", rng.Int63n(100))
	fmt.Printf("Number between 0 and 1000: %d\n", rng.Int63n(1000))
	fmt.Printf("64-bit unsigned integer: %d\n", rng.Uint64())

	// Output: Number between 0 and 10: 4
	// Number between 0 and 100: 67
	// Number between 0 and 1000: 496
	// 64-bit unsigned integer: 16395596082725179435
}
