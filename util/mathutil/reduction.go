package mathutil

// BoundUint32 performs a fast integer reduction,
// For details, see: https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func BoundUint32(x uint32, N uint32) uint32 {
	return uint32((uint64(x) * uint64(N)) >> 32)
}

// BoundUint16 performs a fast integer reduction,
// For details, see: https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func BoundUint16(x uint16, N uint16) uint16 {
	return uint16((uint32(x) * uint32(N)) >> 16)
}
