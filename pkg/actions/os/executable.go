package os

import (
	"os"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionPathExecutables completes executable files from PATH
//
//	nvim
//	chmod
func ActionPathExecutables() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		batch := carapace.Batch()
		manDescriptions := manDescriptions(c)
		dirs := strings.Split(os.Getenv("PATH"), string(os.PathListSeparator))
		for i := len(dirs) - 1; i >= 0; i-- {
			batch = append(batch, actionDirectoryExecutables(dirs[i], c.CallbackValue, manDescriptions))
		}
		return batch.ToA()
	}).Tag("path executables")
}

func actionDirectoryExecutables(dir string, prefix string, manDescriptions map[string]string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if files, err := os.ReadDir(dir); err == nil {
			vals := make([]string, 0)
			for _, f := range files {
				if strings.HasPrefix(f.Name(), prefix) {
					if info, err := f.Info(); err == nil && !f.IsDir() && isExecAny(info.Mode()) {
						description := completers.Description(f.Name())
						if description == "" {
							description = manDescriptions[f.Name()]
						}
						vals = append(vals, f.Name(), description, style.ForPath(dir+"/"+f.Name(), c))
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

func manDescriptions(c carapace.Context) (descriptions map[string]string) {
	descriptions = make(map[string]string)
	if !util.HasPathPrefix(c.CallbackValue) {
		if output, err := c.Command("man", "--names-only", "-k", "^"+c.CallbackValue).Output(); err == nil {
			lines := strings.Split(string(output), "\n")
			r := regexp.MustCompile(`^(?P<name>[^ ]+) [^-]+- (?P<description>.*)$`)
			for _, line := range lines {
				if matches := r.FindStringSubmatch(line); len(matches) > 2 {
					descriptions[matches[1]] = matches[2]
				}
			}
		}
	}
	return
}
