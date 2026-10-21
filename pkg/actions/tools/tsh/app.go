package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionApps completes applications
//
//	grafana (grafana.example.com)
//	prometheus (prometheus.example.com)
func ActionApps() carapace.Action {
	return carapace.ActionExecCommand("tsh", "apps", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var apps []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Spec struct {
				URI        string `json:"uri"`
				PublicAddr string `json:"public_addr"`
			} `json:"spec"`
			Scope string `json:"scope"`
		}

		if err := json.Unmarshal(output, &apps); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, app := range apps {
			desc := app.Spec.PublicAddr
			if desc == "" {
				desc = app.Spec.URI
			}

			name := app.Metadata.Name
			if app.Scope != "" {
				name = app.Scope + "::" + name
			}
			vals = append(vals, name, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
