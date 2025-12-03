package completers

import "testing"

func BenchmarkCompleters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Completers("", false)
	}
}
