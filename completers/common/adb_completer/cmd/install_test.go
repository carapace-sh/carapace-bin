package cmd

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
	"github.com/carapace-sh/carapace/pkg/style"
)

func TestInstallAfterUserFlag(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/common/adb_completer")(func(s *sandbox.Sandbox) {
		s.Files(
			"app.apk", "",
			"notes.txt", "",
		)

		for _, subcommand := range []string{"install", "install-multiple", "install-multi-package"} {
			s.Run(subcommand, "--user", "0", "-r", "").
				Expect(carapace.ActionStyledValues(
					"app.apk", style.ForPath("app.apk", s.NewContext()),
				).NoSpace('/').Tag("files"))
		}
	})
}
