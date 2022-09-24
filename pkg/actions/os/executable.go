package os

import (
	"fmt"
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace/pkg/style"
	"gopkg.in/yaml.v3"
)

// ActionPathExecutables completes executable files from PATH
//
//	nvim
//	chmod
func ActionPathExecutables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()
		dirs := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
		for i := len(dirs) - 1; i >= 0; i-- {
			batch = append(batch, actionDirectoryExecutables(dirs[i], c.CallbackValue))
		}
		return batch.ToA()
	})
}

func actionDirectoryExecutables(dir string, prefix string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if files, err := os.ReadDir(dir); err == nil {
			vals := make([]string, 0)
			for _, f := range files {
				if strings.HasPrefix(f.Name(), prefix) {
					if info, err := f.Info(); err == nil && !f.IsDir() && isExecAny(info.Mode()) {
						if description, err := specDescription(f.Name()); err != nil {
							vals = append(vals, f.Name(), completers.Descriptions[f.Name()], style.ForPath(dir+"/"+f.Name()))
						} else {
							vals = append(vals, f.Name(), description, style.ForPath(dir+"/"+f.Name()))
						}
					}
				}
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		}
		return carapace.ActionValues()
	})
}

func isExecAny(mode os.FileMode) bool {
	return mode&0111 != 0
}

func specDescription(name string) (string, error) {
	confDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(fmt.Sprintf("%v/carapace/specs/%v.yaml", confDir, name))
	if err != nil {
		return "", err
	}
	var s struct {
		Description string
	}

	err = yaml.Unmarshal(content, &s)
	if err != nil {
		return "", err
	}
	return s.Description, nil
}
