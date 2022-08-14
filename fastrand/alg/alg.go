package alg

import (
	"v8.run/go/exp/fastrand/alg/splitmix64"
	"v8.run/go/exp/fastrand/alg/xoshiro256plusplus"
)

func Splitmix64() uint64 {
	return splitmix64.Next()
}

func XoShiro256PlusPlus() uint64 {
	return xoshiro256plusplus.Next()
}
