package randutil

import (
	"crypto/rand"
	"encoding/base64"

	"gopkg.eu.org/exppkgs/fastrand"
)

func RandString(size int) string {
	b := make([]byte, size)
	_, err := rand.Read(b)
	if err != nil {
		// Fallback to a FastRandReader.
		rng := fastrand.AcquireRNG()
		fr := fastrand.FastRandReader{RNG: rng}
		fr.Read(b)
		rng.Release()
	}
	return base64.RawURLEncoding.EncodeToString(b)
}
