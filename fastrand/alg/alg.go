package alg

import (
	"gopkg.eu.org/exppkgs/fastrand/alg/splitmix64"
	"gopkg.eu.org/exppkgs/fastrand/alg/xoshiro256plusplus"
)

func Splitmix64() uint64 {
	return splitmix64.Next()
}

func XoShiro256PlusPlus() uint64 {
	return xoshiro256plusplus.Next()
}
