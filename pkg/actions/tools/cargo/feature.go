package cargo

import (
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionFeatures completes features
//
//	extra (default)
//	stable (default)
func ActionFeatures(path string) carapace.Action {
	return readMetadataAction(path, func(m metadataJson, args []string) carapace.Action {
		vals := make([]string, 0)
		for _, pkg := range m.Packages {
			for name, packages := range pkg.Features {
				vals = append(vals, name, strings.Join(packages, ", "))
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
