/******************************************************************************

Original C code: https://xorshift.di.unimi.it/splitmix64.c
Written in 2015 by Sebastiano Vigna (vigna@acm.org)

To the extent possible under law, the author has dedicated all copyright
and related and neighboring rights to this software to the public domain
worldwide. This software is distributed without any warranty.

See <http://creativecommons.org/publicdomain/zero/1.0/>.

*******************************************************************************/

package splitmix64

import "sync/atomic"

var x uint64 = 11920197720133568065

const IncrementConstant = 0x9e3779b97f4a7c15

func next(x0 uint64) uint64 {
	x0 = (x0 ^ (x0 >> 30)) * 0xbf58476d1ce4e5b9
	x0 = (x0 ^ (x0 >> 27)) * 0x94d049bb133111eb
	return x0 ^ (x0 >> 31)
}

func Next() uint64 {
	n := atomic.AddUint64(&x, IncrementConstant)
	return next(n)
}

func Splitmix64(state *uint64) uint64 {
	*state += IncrementConstant
	return next(*state)
}
