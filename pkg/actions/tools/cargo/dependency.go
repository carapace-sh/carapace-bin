package cargo

import (
	"fmt"
	"strings"

	"github.com/rsteube/carapace"
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
	return readManifestAction(opts.Path, func(m manifestJson, args []string) carapace.Action {
		vals := make([]string, len(m.Dependencies)*2)
		for index, dependency := range m.Dependencies {
			if opts.IncludeVersion {
				vals[index*2] = fmt.Sprintf("%v:%v", dependency.Name, strings.TrimLeft(dependency.Req, "^"))
			} else {
				vals[index*2] = dependency.Name
			}
			vals[(index*2)+1] = dependency.Req
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
