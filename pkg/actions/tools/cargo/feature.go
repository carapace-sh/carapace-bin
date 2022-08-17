package cargo

import (
	"strings"

	"github.com/rsteube/carapace"
)

// ActionFeatures completes features
//
//	extra (default)
//	stable (default)
func ActionFeatures(path string) carapace.Action {
	return readManifestAction(path, func(m manifestJson, args []string) carapace.Action {
		vals := make([]string, 0)
		for name, packages := range m.Features {
			vals = append(vals, name, strings.Join(packages, ", "))
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
