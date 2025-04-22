package tsh

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

func ActionKubeClusters() carapace.Action {
	return carapace.ActionExecCommand("tsh", "kube", "ls", "--format", "json")(func(output []byte) carapace.Action {
		var clusters []struct {
			KubeClusterName string `json:"kube_cluster_name"`
		}

		if err := json.Unmarshal(output, &clusters); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, cluster := range clusters {
			s := style.Green
			vals = append(vals, cluster.KubeClusterName, s)
		}
		return carapace.ActionStyledValues(vals...)
	})
}
