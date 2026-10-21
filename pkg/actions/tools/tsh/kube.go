package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionKubeClusters completes kubernetes clusters
//
//	production (selected)
//	staging
func ActionKubeClusters() carapace.Action {
	return carapace.ActionExecCommand("tsh", "kube", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var clusters []struct {
			KubeClusterName string `json:"kube_cluster_name"`
			Scope           string `json:"scope"`
			Selected        bool   `json:"selected"`
		}

		if err := json.Unmarshal(output, &clusters); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, cluster := range clusters {
			desc := cluster.Scope
			if cluster.Selected {
				if desc != "" {
					desc += ", "
				}
				desc += "selected"
			}
			vals = append(vals, cluster.KubeClusterName, desc, style.Green)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
