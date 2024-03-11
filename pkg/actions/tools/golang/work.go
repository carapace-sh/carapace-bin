package golang

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/util"
)

type work struct {
	Go  string
	Use []struct {
		DiskPath string
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
}

func actionWork(path string, f func(w work) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"work", "edit", "-json"}
		if path != "" {
			args = append(args, path)
		}
		return carapace.ActionExecCommand("go", args...)(func(output []byte) carapace.Action {
			var w work
			if err := json.Unmarshal(output, &w); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(w)
		})
	})
}

// ActionWorkUses completes workspace uses
//
//	./carapace
//	./carapace-bin
func ActionWorkUses(path string) carapace.Action {
	return actionWork(path, func(w work) carapace.Action {
		vals := make([]string, 0)
		for _, use := range w.Use {
			vals = append(vals, use.DiskPath)
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionWorkReplacements completes workspace replacements
// github.com/carapace-sh/carapace-spec (github.com/carapace-sh/carapace-spec@v0.3.0
// github.com/spf13/pflag (../carapace-pflag/)
func ActionWorkReplacements(path string) carapace.Action {
	return actionWork(path, func(w work) carapace.Action {
		vals := make([]string, 0)
		for _, replace := range w.Replace {
			if replace.New.Version != "" {
				vals = append(vals, replace.Old.Path, fmt.Sprintf("%v@%v", replace.New.Path, replace.New.Version))
			} else {
				vals = append(vals, replace.Old.Path, replace.New.Path)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionWorkModules completes workspace modules
//
//	github.com/pelletier/go-toml
//	github.com/carapace-sh/carapace
func ActionWorkModules(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return actionWork(path, func(w work) carapace.Action {
			if path == "" {
				var err error
				if path, err = util.FindReverse(path, "go.work"); err != nil {
					return carapace.ActionMessage(err.Error())
				}
			}
			abs, err := c.Abs(path)
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			batch := carapace.Batch()
			for _, use := range w.Use {
				dir := fmt.Sprintf("%v/%v", filepath.Dir(abs), use.DiskPath)
				batch = append(batch, ActionModules(ModuleOpts{Direct: true, Indirect: true}).Chdir(dir))
			}
			return batch.ToA()
		})
	})
}
