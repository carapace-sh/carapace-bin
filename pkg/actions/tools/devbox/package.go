package devbox

import (
	"encoding/json"
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
)

type devbox struct {
	Packages []string `json:"packages"`
}

// ActionInstalledPackages completes installed packaages
//
//	rustc
//	cargo
func ActionInstalledPackages() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path, err := util.FindReverse(c.Dir, "devbox.json")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var d devbox
		if err := json.Unmarshal(content, &d); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return carapace.ActionValues(d.Packages...)
	})
}
