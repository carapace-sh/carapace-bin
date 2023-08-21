package tsh

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionClusters() carapace.Action {
	return carapace.ActionExecCommand("tsh", "clusters", "--format", "json")(func(output []byte) carapace.Action {
		var clusters []struct {
			ClusterName string `json:"cluster_name"`
			Status      string
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
			vals = append(vals, cluster.ClusterName, s)
		}
		return carapace.ActionStyledValues(vals...)
	})
}
