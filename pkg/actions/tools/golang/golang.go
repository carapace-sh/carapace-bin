// Package golang contains go related actions
package golang

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

// ActionBuildTags completes build tags
//   release
//   debug
func ActionBuildTags() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		root := c.Dir
		if path, err := util.FindReverse(root, "go.mod"); err == nil {
			root = path
		} else if path, err = util.FindReverse(root, ".git"); err == nil {
			root = path
		}

		tags := make(map[string]bool)

		filepath.Walk(filepath.Dir(root), func(path string, info fs.FileInfo, err error) error {
			if !info.IsDir() && strings.HasSuffix(path, ".go") {
				for _, tag := range readTags(path) {
					tags[tag] = true
				}
			}
			return nil
		})

		vals := make([]string, 0, len(tags))
		for tag := range tags {
			vals = append(vals, tag)
		}
		return carapace.ActionValues(vals...)
	})
}

var replacer = strings.NewReplacer(
	" NOT ", " ",
	" AND ", " ",
	" OR ", " ",
	",", " ",
	"|", " ",
	"&", " ",
	"!", " ",
	"(", " ",
	")", " ",
)

func readTags(path string) []string {
	if file, err := os.Open(path); err == nil {
		defer file.Close()

		scanner := bufio.NewScanner(file)
		scanner.Scan()

		if s := scanner.Text(); strings.HasPrefix(s, "//go:build ") {
			s := strings.TrimPrefix(s, "//go:build ")
			s = replacer.Replace(s)

			return strings.Fields(s)
		}
	}
	return []string{}
}
