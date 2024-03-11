package nix

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionNarFileContents completes contents of given nar file
//
//	go.mod
//	cmd/
func ActionNarFileContents(file string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if c.Value == "" {
			c.Value = "/"
		}
		return carapace.ActionExecCommand("nix", "--extra-experimental-features", "nix-command", "nar", "ls", "--long", file, c.Value)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")

			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				fields := strings.Fields(line) // TODO handle filepaths with space in them correctly
				if strings.HasPrefix(fields[0], "d") {
					vals = append(vals, strings.TrimPrefix(fields[2], ".")+"/")
				} else {
					vals = append(vals, strings.TrimPrefix(fields[2], "."))
				}
			}
			return carapace.ActionValues(vals...).StyleF(style.ForPathExt)
		})
	})
}
