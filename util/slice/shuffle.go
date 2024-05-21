package slice

import (
	"gopkg.eu.org/exppkgs/fastrand"
)

func Shuffle[T any](data []T) {
	fastrand.ShuffleSlice(data)
}
