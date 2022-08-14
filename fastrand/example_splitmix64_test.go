package fastrand_test

import (
	"fmt"

	"v8.run/go/exp/fastrand/alg"
)

func Example_splitmix64() {
	fmt.Println(alg.Splitmix64())
	fmt.Println(alg.Splitmix64())
	fmt.Println(alg.Splitmix64())
	fmt.Println(alg.Splitmix64())
	fmt.Println(alg.Splitmix64())

	// Output: 9910846700779238901
	// 5865067318551043833
	// 4341075433095661249
	// 9570429342510947219
	// 10923570614615297618
}
