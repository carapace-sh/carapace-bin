package k3d

import (
	"encoding/json"
	"fmt"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type node struct {
	Name          string
	Role          string
	RuntimeLabels struct {
		K3dCluster string `json:"k3d.cluster"`
	}
	State struct {
		Running bool
	}
}

func (n node) style() string {
	if n.State.Running {
		return style.Carapace.KeywordPositive
	}
	return style.Carapace.KeywordNegative
}

func ActionNodes() carapace.Action {
	return carapace.ActionExecCommand("k3d", "node", "list", "--output", "json")(func(output []byte) carapace.Action {
		var nodes []node
		if err := json.Unmarshal(output, &nodes); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, n := range nodes {
			vals = append(vals, n.Name, fmt.Sprintf("%v.%v", n.RuntimeLabels.K3dCluster, n.Role), n.style())
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
