package os

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace/pkg/style"
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
						vals = append(vals, f.Name(), completers.Description(f.Name()), style.ForPath(dir+"/"+f.Name()))
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
