package golang

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace/pkg/style"
)

type mod struct {
	Module struct {
		Path       string
		Deprecated string
	}
	Go        string
	Toolchain string
	Tool      []struct {
		Path string
	}
	Godebug []struct {
		Key   string
		Value string
	}
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
			Path    string
			Version string
		}
		New struct {
			Path    string
			Version string
		}
	}
	Retract []struct {
		Low       string
		High      string
		Rationale string
	}
	Ignore []struct {
		Path string
	}
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

func actionGoMod(f func(m mod) carapace.Action) carapace.Action {
	return carapace.ActionExecCommand("go", "mod", "edit", "-json")(func(output []byte) carapace.Action {
		var m mod
		if err := json.Unmarshal(output, &m); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(m)
	})
}

// ActionModVersions completes tags of module repositiory
//
//	v0.0.3 (da4ba1e4f4e22b7b278d450719820a08d9e51f79)
//	v0.0.4 (e6815a3c05828b937fce6183f14ba68cc3173726)
func ActionModVersions() carapace.Action {
	return actionGoMod(func(m mod) carapace.Action {
		return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + m.Module.Path, Tags: true})
	}).Tag("mod versions")
}

// ActionModGodebugs completes godebug instructions
// TODO needs https://github.com/golang/go/issues/75105
func ActionModGodebugs() carapace.Action {
	return actionGoMod(func(m mod) carapace.Action {
		vals := make([]string, 0)
		for _, d := range m.Godebug {
			vals = append(vals, d.Key, d.Value)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("mod godebugs")
}

// ActionModRetracts completes retract instructions
//
//	[v0.5.0,v0.5.4]
//	v0.4.0
func ActionModRetracts() carapace.Action {
	return actionGoMod(func(m mod) carapace.Action {
		vals := make([]string, 0)
		for _, r := range m.Retract {
			switch r.Low {
			case r.High:
				vals = append(vals, r.Low, r.Rationale)
			default:
				vals = append(vals, fmt.Sprintf("[%v,%v]", r.Low, r.High), r.Rationale)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("mod retracts")
}

// ActionModIgnores completes ignore instructions
//
//	dist
//	./third_party
func ActionModIgnores() carapace.Action {
	return actionGoMod(func(m mod) carapace.Action {
		vals := make([]string, 0)
		for _, i := range m.Ignore {
			vals = append(vals, i.Path)
		}
		return carapace.ActionValues(vals...)
	}).Tag("mod ignores")
}

// ActionModTools completes tool instructions
//
//	github.com/some/tool
//	github.com/another/tool
func ActionModTools() carapace.Action {
	return actionGoMod(func(m mod) carapace.Action {
		vals := make([]string, 0)
		for _, t := range m.Tool {
			vals = append(vals, t.Path)
		}
		return carapace.ActionValues(vals...)
	}).Tag("mod tools")
}

// ActionModules completes modules
//
//	github.com/carapace-sh/carapace
//	github.com/carapace-sh/carapace-spec@v0.0.1
func ActionModules(opts ModuleOpts) carapace.Action {
	return actionGoMod(func(m mod) carapace.Action {
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
