package cmd

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
)

func TestMacros(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		s.Run("--macro", "color.XtermColorNames", "G").
			Expect(carapace.ActionStyledValues(
				"Green", "color2",
				"Grey", "color8",
			).Tag("xterm colors").
				Usage("--macro [macro] [...args]"))
	})
}
