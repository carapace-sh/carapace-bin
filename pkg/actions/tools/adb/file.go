package adb

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionFiles completes device files
func ActionFiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		switch c.Value {
		case "":
			c.Value = "/"
		default:
			c.Value = filepath.Dir(c.Value)
		}

		return carapace.Batch(
			actionDirectories(),
			actionFiles(),
		).Invoke(c).Merge().ToA()
	}).MultiParts("/").StyleF(style.ForPathExt).Tag("device files")
}

func actionDirectories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("adb", "shell", fmt.Sprintf("find -L %#v -maxdepth 1 -type d 2>/dev/null ; true", c.Value))(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if line != "/" && line != c.Value {
					vals = append(vals, line+"/")
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}

func actionFiles() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("adb", "shell", fmt.Sprintf("find -L %#v -maxdepth 1 -type f -o -type l 2>/dev/null ; true", c.Value))(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)
			for _, line := range lines[:len(lines)-1] {
				if line != c.Value {
					vals = append(vals, line)
				}
			}
			return carapace.ActionValues(vals...)
		})
	})
}
