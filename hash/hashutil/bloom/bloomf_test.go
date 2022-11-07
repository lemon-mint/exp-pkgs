package bloom

import (
	"testing"
)

func Test(t *testing.T) {
	t.Run("Bloom", func(t *testing.T) {
		var list = []string{
			"apple",
			"banana",
			"cherry",
		}
		bf := NewBloom(uint64(len(list)), 0.01)
		for _, s := range list {
			bf.SetString(s)
		}

		for _, s := range list {
			if !bf.GetString(s) {
				t.Errorf("TestString(%q) = false, want true", s)
			}
		}
	})
}
