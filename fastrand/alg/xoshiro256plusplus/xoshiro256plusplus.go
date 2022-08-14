/******************************************************************************

Original C code: https://prng.di.unimi.it/xoshiro256plusplus.c
Written in 2019 by David Blackman and Sebastiano Vigna (vigna@acm.org)

To the extent possible under law, the author has dedicated all copyright
and related and neighboring rights to this software to the public domain
worldwide. This software is distributed without any warranty.

See <http://creativecommons.org/publicdomain/zero/1.0/>.

*******************************************************************************/

package xoshiro256plusplus

import (
	"sync"

	"v8.run/go/exp/fastrand/alg/splitmix64"
)

type State [4]uint64

func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

func (s *State) Next() uint64 {
	var result uint64 = rotl(s[0]+s[3], 23) + s[0]
	var t uint64 = s[1] << 17

	s[2] ^= s[0]
	s[3] ^= s[1]
	s[1] ^= s[2]
	s[0] ^= s[3]

	s[2] ^= t

	s[3] = rotl(s[3], 45)

	return result
}

func (s *State) init(seed uint64) {
	seed = splitmix64.Splitmix64(&seed)
	s[0] = splitmix64.Splitmix64(&seed)
	s[1] = splitmix64.Splitmix64(&seed)
	s[2] = splitmix64.Splitmix64(&seed)
	s[3] = splitmix64.Splitmix64(&seed)
}

// NewState initializes xoshiro256++ state.
func NewState() State {
	var s State
	s.init(splitmix64.Next())
	return s
}

// NewStateWithSeed initializes xoshiro256++ state with a seed.
func NewStateWithSeed(seed uint64) State {
	var s State
	s.init(seed)
	return s
}

const seed uint64 = 9674114761913717981

var gState State = NewStateWithSeed(seed)
var mu sync.Mutex

func Next() uint64 {
	mu.Lock()
	v := gState.Next()
	mu.Unlock()
	return v
}
