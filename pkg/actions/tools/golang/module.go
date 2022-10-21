package golang

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/rsteube/carapace/pkg/style"
)

type module struct {
	Path     string
	Version  string
	Indirect bool
	Replace  *struct {
		Path    string
		Version string
	}
	Exclude *struct {
		Path    string
		Version string
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

// ActionModules completes ModuleOpts
//
//	github.com/rsteube/carapace
//	github.com/rsteube/carapace-spec@v0.0.1
func ActionModules(opts ModuleOpts) carapace.Action {
	return carapace.ActionExecCommand("go", "list", "-m", "-json", "all")(func(output []byte) carapace.Action {
		dec := json.NewDecoder(bytes.NewReader(output))

		vals := make([]string, 0)
		for {
			var m module
			if err := dec.Decode(&m); err == io.EOF {
				break
			} else if err != nil {
				return carapace.ActionMessage(err.Error())

			}
			if opts.Direct && !m.Indirect && m.Replace == nil { // TODO Exclude
				vals = append(vals, formatModule(m.Path, m.Version, opts.IncludeVersion), style.Blue)
			} else if opts.Indirect && m.Indirect && m.Replace == nil { // TODO Exclude
				vals = append(vals, formatModule(m.Path, m.Version, opts.IncludeVersion), style.Gray)
			} else if opts.Replace && m.Replace != nil { // TODO Exclude
				vals = append(vals, formatModule(m.Path, m.Version, opts.IncludeVersion), style.Magenta)
			}
		}

		batch := carapace.Batch(
			carapace.ActionStyledValues(vals...),
		)
		if opts.Exclude {
			batch = append(batch, actionModuleExcludes(opts.IncludeVersion))
		}

		return batch.ToA()
	})
}

func actionModuleExcludes(includeVersion bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "go.mod")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, line := range strings.Split(string(content), "\n") {
			if strings.HasPrefix(line, "exclude ") {
				if fields := strings.Fields(line); len(fields) == 3 {
					vals = append(vals, formatModule(fields[1], fields[2], includeVersion), style.Red)
				}
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
