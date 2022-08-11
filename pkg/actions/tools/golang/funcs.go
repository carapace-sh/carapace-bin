package golang

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

func readFuncs(path string) []string {
	// TODO optimize performance
	if file, err := os.Open(path); err == nil {
		defer file.Close()

		r := regexp.MustCompile(`^func (\([^)]+\) )?(?P<func>[^([]+).*$`)
		scanner := bufio.NewScanner(file)

		names := make([]string, 0)
		for scanner.Scan() {
			if s := scanner.Text(); strings.HasPrefix(s, "func ") {
				if matches := r.FindStringSubmatch(s); matches != nil {
					names = append(names, matches[2])
				}
			}
		}
		return names
	}
	return []string{}
}

// ActionFuncs completes go funcs
//   main
//   ActionFuncs
func ActionFuncs(files ...string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		root := c.Dir
		if path, err := util.FindReverse(root, "go.mod"); err == nil {
			root = path
		} else if path, err = util.FindReverse(root, ".git"); err == nil {
			root = path
		}

		names := make(map[string]bool)

		// TODO optimize performance - goroutines?
		if len(files) == 0 {
			filepath.Walk(filepath.Dir(root), func(path string, info fs.FileInfo, err error) error {
				if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
					for _, name := range readFuncs(path) {
						names[name] = true
					}
				}
				return nil
			})
		} else {
			for _, file := range files {
				for _, name := range readFuncs(file) {
					names[name] = true
				}
			}
		}

		vals := make([]string, 0, len(names))
		for name := range names {
			vals = append(vals, name)
		}
		return carapace.ActionValues(vals...)
	})
}
