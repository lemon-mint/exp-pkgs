package alg

import "github.com/lemon-mint/experiment/fastrand/alg/splitmix64"

func Splitmix64() uint64 {
	return splitmix64.Next()
}
