package completers

import (
	"testing"

	"github.com/carapace-sh/carapace-bridge/pkg/choices"
)

func BenchmarkCompleters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Completers(choices.Choice{}, false)
	}
}
