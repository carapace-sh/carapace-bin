package action

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

type plugin struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Source  string `json:"source"`
	Enabled bool   `json:"enabled"`
}

type pluginListResult struct {
	Plugins []plugin `json:"plugins"`
	Type    string   `json:"type"`
}

func ActionPlugins() carapace.Action {
	return carapace.ActionExecCommand("herdr", "plugin", "list", "--json")(func(output []byte) carapace.Action {
		wrapper := struct {
			Result pluginListResult `json:"result"`
		}{}
		if err := json.Unmarshal(output, &wrapper); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, p := range wrapper.Result.Plugins {
			desc := p.Name
			if !p.Enabled {
				desc += " (disabled)"
			}
			vals = append(vals, p.ID, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Tag("plugins")
}
