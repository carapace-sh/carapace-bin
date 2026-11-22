package cmd

import (
	"testing"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/sandbox"
	"github.com/carapace-sh/carapace/pkg/style"
)

func TestSideload(t *testing.T) {
	sandbox.Package(t, "github.com/carapace-sh/carapace-bin/completers/common/adb_completer")(func(s *sandbox.Sandbox) {
		s.Files(
			"firmware.bin", "",
			"notes.txt", "",
			"ota.zip", "",
		)

		s.Run("sideload", "").
			Expect(carapace.ActionStyledValues(
				"firmware.bin", style.ForPath("firmware.bin", s.NewContext()),
				"ota.zip", style.ForPath("ota.zip", s.NewContext()),
			).NoSpace('/').Tag("files"))
	})
}
