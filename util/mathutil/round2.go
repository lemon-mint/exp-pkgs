/******************************************************************************

Original C code: https://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2

*******************************************************************************/

package mathutil

func RoundUpPowerOf2(v uint32) uint32 {
	v--
	v |= v >> 1
	v |= v >> 2
	v |= v >> 4
	v |= v >> 8
	v |= v >> 16
	v++
	return v
}
