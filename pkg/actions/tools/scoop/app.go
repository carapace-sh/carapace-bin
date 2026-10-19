// package scoop contains Scoop related actions
package scoop

import (
	"encoding/json"
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
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "scoop export")(func(output []byte) carapace.Action {
		var export scoopExport
		if err := json.Unmarshal(output, &export); err != nil {
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
	return carapace.ActionExecCommand("powershell", "-NoProfile", "-Command", "$d = if ($env:SCOOP) { $env:SCOOP } else { \"$env:USERPROFILE\\scoop\" }; Get-ChildItem -Path \"$d\\buckets\\*\\bucket\\*.json\" | Select-Object -ExpandProperty BaseName | ConvertTo-Json -Compress")(func(output []byte) carapace.Action {
		output = []byte(strings.TrimSpace(string(output)))
		if len(output) == 0 || string(output) == "null" {
			return carapace.ActionValues()
		}
		var apps []string
		if err := json.Unmarshal(output, &apps); err != nil {
			var app string
			if err2 := json.Unmarshal(output, &app); err2 != nil {
				return carapace.ActionMessage(err.Error())
			}
			apps = []string{app}
		}
		return carapace.ActionValues(apps...).Unique().Tag("apps")
	})
}
