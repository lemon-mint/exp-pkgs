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
}
