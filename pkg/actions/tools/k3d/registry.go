package k3d

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type registry struct {
	Name          string
	RuntimeLabels struct {
		K3dCluster string `json:"k3d.cluster"`
	}
	State struct {
		Running bool
	}
}

func (r registry) style() string {
	if r.State.Running {
		return style.Carapace.KeywordPositive
	}
	return style.Carapace.KeywordNegative
}

// ActionRegistries completes k3d registries
func ActionRegistries() carapace.Action {
	return carapace.ActionExecCommand("k3d", "registry", "list", "--output", "json")(func(output []byte) carapace.Action {
		var registries []registry
		if err := json.Unmarshal(output, &registries); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, r := range registries {
			vals = append(vals, r.Name, fmt.Sprintf("%v", r.RuntimeLabels.K3dCluster), r.style())
		}
		return carapace.ActionStyledValuesDescribed(vals...)
	})
}
