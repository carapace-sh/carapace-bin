// package scoop contains Scoop related actions
package scoop

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
)

// scoopExport is the JSON structure output by `scoop export`.
type scoopExport struct {
	Apps    []scoopApp    `json:"apps"`
	Buckets []scoopBucket `json:"buckets"`
}

type scoopApp struct {
	Name    string `json:"Name"`
	Version string `json:"Version"`
	Source  string `json:"Source"`
}

// ActionInstalledApps completes installed apps with their version as description
//
//	bat (0.24.0)
//	git (2.45.0)
func ActionInstalledApps() carapace.Action {
	return carapace.ActionExecCommand("scoop", "export")(func(output []byte) carapace.Action {
		var export scoopExport
		if err := json.Unmarshal(bytes.TrimPrefix(output, []byte{0xEF, 0xBB, 0xBF}), &export); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, app := range export.Apps {
			vals = append(vals, app.Name, app.Version)
		}
		return carapace.ActionValuesDescribed(vals...).Tag("apps")
	})
}

// ActionAvailableApps completes app names available in local buckets
//
//	bat
//	git
func ActionAvailableApps() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		root := loadConfig().scoopDir()
		if root == "" {
			return carapace.ActionValues()
		}
		matches, err := filepath.Glob(filepath.Join(root, "buckets", "*", "bucket", "*.json"))
		if err != nil {
			return carapace.ActionValues()
		}
		vals := make([]string, 0, len(matches))
		for _, match := range matches {
			vals = append(vals, strings.TrimSuffix(filepath.Base(match), filepath.Ext(match)))
		}
		return carapace.ActionValues(vals...).Unique().Tag("apps")
	})
}
