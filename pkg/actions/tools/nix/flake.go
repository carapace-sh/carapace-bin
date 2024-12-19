package nix

import (
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionFlakes completes flakes currently available
//
// nixpkgs
// .
func ActionFlakes() carapace.Action {
	return carapace.ActionExecCommand("nix", "registry", "list")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")

		vals := make([]string, 0)
		// `nix registry list` outputs a blank line after the listings
		for _, line := range lines[:len(lines)-1] {
			fields := strings.Fields(line)
			name := strings.Split(fields[1], ":")
			if name[0] == "flake" {
				vals = append(vals, name[1], fields[2], styleForRegistry(fields[0]))
			}
		}
		return carapace.Batch(
			carapace.ActionValues(".", "..").StyleF(style.ForPath),
			carapace.ActionDirectories(),
			carapace.ActionStyledValuesDescribed(vals...),
		).ToA()
	}).Cache(time.Minute)
}

func styleForRegistry(s string) string {
	switch s {
	case "global":
		return style.Blue
	case "system":
		return style.Yellow
	default:
		return style.Default
	}
}
