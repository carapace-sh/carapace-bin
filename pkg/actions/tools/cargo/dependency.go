package cargo

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

type DependencyOpts struct {
	Path           string
	IncludeVersion bool
}

// ActionDependencies completes dependencies
//
//	assert_cmd (^2.0.2)
//	chrono (^0.4.21)
func ActionDependencies(opts DependencyOpts) carapace.Action {
	return readMetadataAction(opts.Path, func(m metadataJson, args []string) carapace.Action {
		vals := make([]string, 0)
		for _, pkg := range m.Packages {
			for _, dependency := range pkg.Dependencies {
				if opts.IncludeVersion {
					vals = append(vals, fmt.Sprintf("%v:%v", dependency.Name, strings.TrimLeft(dependency.Req, "^")), dependency.Req)
				} else {
					vals = append(vals, dependency.Name, dependency.Req)
				}
			}
		}
		return carapace.ActionValuesDescribed(vals...).Unique()
	})
}
