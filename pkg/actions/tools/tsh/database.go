package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionDatabases completes databases
//
//	postgres (postgres://postgres.example.com:5432)
//	mysql (mysql://mysql.example.com:3306)
func ActionDatabases() carapace.Action {
	return carapace.ActionExecCommand("tsh", "db", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var databases []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Spec struct {
				Protocol string `json:"protocol"`
				URI      string `json:"uri"`
			} `json:"spec"`
		}

		if err := json.Unmarshal(output, &databases); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, database := range databases {
			desc := database.Spec.Protocol
			if database.Spec.URI != "" {
				desc = database.Spec.URI
			}
			vals = append(vals, database.Metadata.Name, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
