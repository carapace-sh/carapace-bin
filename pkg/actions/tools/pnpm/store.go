package pnpm

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionStorePackages completes packages in the store
func ActionStorePackages() carapace.Action {
	return carapace.ActionExecCommand("pnpm", "store", "status")(func(output []byte) carapace.Action {
		lines := strings.Split(strings.TrimSpace(string(output)), "\n")
		vals := make([]string, 0)
		for _, line := range lines {
			if line != "" && !strings.Contains(line, "Packages:") && !strings.Contains(line, "Size:") {
				vals = append(vals, line)
			}
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionStorePrune completes store prune options
func ActionStorePrune() carapace.Action {
	return carapace.ActionValues("--verify-store-integrity")
}

// ActionStoreAdd completes store add options
func ActionStoreAdd() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			return ActionPackageNames("").NoSpace()
		case 1:
			return ActionPackageVersions(PackageOpts{Package: c.Parts[0]})
		default:
			return carapace.ActionValues()
		}
	})
}
