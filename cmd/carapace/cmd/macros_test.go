package cmd

import (
	"testing"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/sandbox"
)

func TestMacros(t *testing.T) {
	sandbox.Package(t, "github.com/rsteube/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		s.Run("--macros", "color.XtermColorNames", "G").
			Expect(carapace.ActionStyledValues(
				"Green", "color2",
				"Grey", "color8",
			).Usage("--macros [macro] ..."))
	})
}
