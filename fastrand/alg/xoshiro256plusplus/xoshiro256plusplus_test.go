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
	"testing"
)

func BenchmarkXoShiro256PlusPlus(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		s := NewState()
		for p.Next() {
			s.Next()
		}
	})
}
