package slice

import (
	"v8.run/go/exp/fastrand"
)

func Shuffle[T any](data []T) {
	fastrand.ShuffleSlice(data)
}
