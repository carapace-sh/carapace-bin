package cmd

import (
	"testing"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/sandbox"
	"github.com/rsteube/carapace/pkg/style"
)

func TestInvokeFlags(t *testing.T) {
	sandbox.Package(t, "github.com/rsteube/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		s.Run("tail", "export", "tail", "--fo").
			Expect(carapace.ActionStyledValuesDescribed(
				"--follow", "output appended data as the file grows", style.Yellow,
			).Tag("flags").
				NoSpace('.').
				Usage("carapace [flags] [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]")) // TODO fix usage

		s.Run("tail", "export", "tail", "--follow=").
			Expect(carapace.ActionValues(
				"name",
				"descriptor",
			).Prefix("--follow=").
				Usage("output appended data as the file grows"))
	})
}

func TestInvokePositional(t *testing.T) {
	sandbox.Package(t, "github.com/rsteube/carapace-bin/cmd/carapace")(func(s *sandbox.Sandbox) {
		s.Run("git", "export", "git", "checko").
			Expect(carapace.Batch(
				carapace.ActionValuesDescribed(
					"checkout", "Switch branches or restore working tree files",
				).Style(style.Blue).
					Tag("main commands"),
				carapace.ActionValuesDescribed(
					"checkout-index", "Copy files from the index to the working tree",
				).Style(style.Of(style.Dim, style.Yellow)).
					Tag("low-level manipulator commands"),
			).ToA().
				Usage("carapace [flags] [COMPLETER] [bash|elvish|fish|nushell|oil|powershell|tcsh|xonsh|zsh]")) // TODO fix usage
	})
}
