package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

// ActionClusters completes teleport clusters
//
//	production (root, selected)
//	staging (leaf, offline)
func ActionClusters() carapace.Action {
	return carapace.ActionExecCommand("tsh", "clusters", "--format", "json")(func(output []byte) carapace.Action {
		var clusters []struct {
			ClusterName string `json:"cluster_name"`
			Status      string `json:"status"`
			ClusterType string `json:"cluster_type"`
			Selected    bool   `json:"selected"`
		}

		if err := json.Unmarshal(output, &clusters); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, cluster := range clusters {
			s := style.Green
			if cluster.Status != "online" {
				s = style.Red
			}

			desc := cluster.ClusterType
			if cluster.Selected {
				desc += ", selected"
			}
			if cluster.Status != "" {
				desc += ", " + cluster.Status
			}
			vals = append(vals, cluster.ClusterName, desc, s)
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
