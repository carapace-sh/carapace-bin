package completers

import (
	"testing"

	"github.com/carapace-sh/carapace-bridge/pkg/choice"
)

func BenchmarkCompleters(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Completers(choice.Choice{}, false)
	}
}
