package nix

import (
	"fmt"
	"runtime"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionSystems completes nix systems.
// The system of the current machine is highlighted.
//
//	aarch64-darwin
//	x86_64-linux
func ActionSystems() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		currentSystem, err := getCurrentSystem(c)
		if err != nil {
			currentSystem = fmt.Sprintf("%v-%v", currentArch(), runtime.GOOS)
		}

		return carapace.ActionValues(
			"aarch64-darwin",
			"aarch64-linux",
			"armv6l-linux",
			"armv7l-linux",
			"i686-linux",
			"powerpc64le-linux",
			"riscv64-linux",
			"x86_64-darwin",
			"x86_64-linux",
		).StyleF(func(s string, sc style.Context) string {
			if s == currentSystem {
				return style.Blue
			}
			return style.Default
		})
	}).Tag("systems")
}

// currentArch translates GOARCH to the nix equivalent
func currentArch() string {
	switch runtime.GOARCH {
	case "386":
		return "i686"
	case "amd64":
		return "x86_64"
	case "arm64":
		return "aarch64"
	case "ppc64le":
		return "powerpc64le"
	default:
		return runtime.GOARCH
	}
}
