package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionMCPServers completes MCP servers
//
//	grafana-mcp (grafana.example.com)
//	github-mcp (github.example.com)
func ActionMCPServers() carapace.Action {
	return carapace.ActionExecCommand("tsh", "mcp", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var servers []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Spec struct {
				URI        string `json:"uri"`
				PublicAddr string `json:"public_addr"`
			} `json:"spec"`
			Scope string `json:"scope"`
		}

		if err := json.Unmarshal(output, &servers); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, server := range servers {
			desc := server.Spec.PublicAddr
			if desc == "" {
				desc = server.Spec.URI
			}

			name := server.Metadata.Name
			if server.Scope != "" {
				name = server.Scope + "::" + name
			}
			vals = append(vals, name, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
