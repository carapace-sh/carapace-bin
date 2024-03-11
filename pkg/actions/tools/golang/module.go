package golang

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type mod struct {
	Module struct {
		Path string
	}
	Go      string
	Require []struct {
		Path     string
		Version  string
		Indirect bool
	}
	Exclude []struct {
		Path    string
		Version string
	}
	Replace []struct {
		Old struct {
			Path string
		}
		New struct {
			Path    string
			Version string
		}
	}
	Retract interface{}
}

type ModuleOpts struct {
	Direct         bool
	Indirect       bool
	Replace        bool
	Exclude        bool
	IncludeVersion bool
}

func (o ModuleOpts) Default() ModuleOpts {
	o.Direct = true
	o.Indirect = true
	return o
}

// ActionModules completes ModuleOpts
//
//	github.com/carapace-sh/carapace
//	github.com/carapace-sh/carapace-spec@v0.0.1
func ActionModules(opts ModuleOpts) carapace.Action {
	return carapace.ActionExecCommand("go", "mod", "edit", "-json")(func(output []byte) carapace.Action {
		var m mod
		if err := json.Unmarshal(output, &m); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		if opts.Direct && m.Require != nil {
			for _, r := range m.Require {
				if opts.Direct && !r.Indirect {
					vals = append(vals, formatModule(r.Path, r.Version, opts.IncludeVersion), style.Blue)
				} else if opts.Indirect && r.Indirect {
					vals = append(vals, formatModule(r.Path, r.Version, opts.IncludeVersion), style.Of(style.Dim, style.White))
				}
			}
		}

		if opts.Replace && m.Replace != nil {
			for _, r := range m.Replace {
				vals = append(vals, r.Old.Path, style.Magenta)
			}
		}

		if opts.Exclude && m.Exclude != nil {
			for _, r := range m.Exclude {
				vals = append(vals, formatModule(r.Path, r.Version, opts.IncludeVersion), style.Red)
			}

		}
		return carapace.ActionStyledValues(vals...)
	})
}

func formatModule(path, version string, includeVersion bool) string {
	if includeVersion {
		return fmt.Sprintf("%v@%v", path, version)
	}
	return path
}

// ActionModuleDownloadModes completes module download modes
//
//	readonly (disable implicit automatic updating of go.mod)
//	vendor (assumes that the vendor directory holds the correct copies of dependencies)
func ActionModuleDownloadModes() carapace.Action {
	return carapace.ActionValuesDescribed(
		"readonly", "disable implicit automatic updating of go.mod",
		"mod", "", // TODO description
		"vendor", "assumes that the vendor directory holds the correct copies of dependencies",
	)
}
