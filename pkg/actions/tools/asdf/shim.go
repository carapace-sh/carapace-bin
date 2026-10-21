package asdf

import (
	"os"
	"path/filepath"

	"github.com/carapace-sh/carapace"
)

// ActionShims completes shims
//
//	ruby
//	node
//	python
func ActionShims() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		entries, err := os.ReadDir(filepath.Join(dataDir(), "shims"))
		if err != nil {
			return carapace.ActionValues()
		}
		var vals []string
		for _, entry := range entries {
			if !entry.IsDir() {
				vals = append(vals, entry.Name())
			}
		}
		return carapace.ActionValues(vals...).Tag("shims")
	})
}
