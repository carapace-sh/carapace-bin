package pacman

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionPackageFiles completes package files
//
//	/usr/share/man/man1/git-clone.1.gz (git)
//	/usr/bin/tig (tig)
func ActionPackageFiles(packages ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		files := make([]string, 0)
		names := make([]string, 0)
		for _, pkg := range packages {
			if strings.Contains(pkg, ".") { // local file
				files = append(files, pkg)
			} else {
				names = append(names, pkg)
			}
		}
		return carapace.Batch(
			actionPackageFiles(true, files...),
			actionPackageFiles(false, names...),
		).ToA()
	})
}

func actionPackageFiles(files bool, packages ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(packages) == 0 {
			return carapace.ActionValues()
		}

		opts := []string{"-Q", "-l"}
		if files {
			opts = append(opts, "-p")
		}

		return carapace.ActionExecCommand("pacman", append(opts, packages...)...)(func(output []byte) carapace.Action {
			lines := strings.Split(string(output), "\n")
			vals := make([]string, 0)

			for _, line := range lines[:len(lines)-1] {
				if splitted := strings.SplitN(line, " ", 2); len(splitted) == 2 {
					vals = append(vals, splitted[1], splitted[0])
				}
			}
			return carapace.ActionValuesDescribed(vals...).StyleF(style.ForPathExt).Invoke(c).ToMultiPartsA("/")
		})
	})
}
