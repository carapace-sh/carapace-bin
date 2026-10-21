package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
)

// ActionNodes completes SSH nodes
//
//	node1.example.com (10.0.0.1:3022)
//	node2.example.com (10.0.0.2:3022)
func ActionNodes() carapace.Action {
	return carapace.ActionExecCommand("tsh", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var nodes []struct {
			Metadata struct {
				Name string `json:"name"`
			} `json:"metadata"`
			Spec struct {
				Addr     string `json:"addr"`
				Hostname string `json:"hostname"`
			} `json:"spec"`
		}

		if err := json.Unmarshal(output, &nodes); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, node := range nodes {
			value := node.Spec.Hostname
			if value == "" {
				value = node.Metadata.Name
			}

			desc := node.Spec.Addr
			if desc == "" {
				desc = node.Metadata.Name
			}
			vals = append(vals, value, desc)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
