package slice

import (
	"github.com/lemon-mint/experiment/fastrand"
)

func Shuffle[T any](data []T) {
	fastrand.ShuffleSlice(data)
}
