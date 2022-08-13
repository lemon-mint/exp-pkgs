package alg

import (
	"github.com/lemon-mint/experiment/fastrand/alg/splitmix64"
	"github.com/lemon-mint/experiment/fastrand/alg/xoshiro256plusplus"
)

func Splitmix64() uint64 {
	return splitmix64.Next()
}

func XoShiro256PlusPlus() uint64 {
	return xoshiro256plusplus.Next()
}
