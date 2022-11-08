package bloom

import (
	"testing"
)

func TestBloom(t *testing.T) {
	t.Run("bloom-3", func(t *testing.T) {
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

func BenchmarkBloom(b *testing.B) {
	var list = []string{
		"apple",
		"banana",
		"cherry",
		"golang",
		"python",
		"ruby",
		"rust",
		"java",
	}
	bf := NewBloom(uint64(len(list)), 0.01)
	for _, s := range list {
		bf.SetString(s)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bf.GetString("apple")
	}
}
