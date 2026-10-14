package brew

import (
	"encoding/json"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type service struct {
	Name   string
	Status string
}

// ActionServices completes services
//
//	postgresql@17 (started)
//	unbound (none)
func ActionServices() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		c.Setenv("HOMEBREW_NO_AUTO_UPDATE", "1")
		return carapace.ActionExecCommand("brew", "services", "list", "--json")(func(output []byte) carapace.Action {
			var services []service
			if err := json.Unmarshal(output, &services); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, service := range services {
				vals = append(vals, service.Name, service.Status, style.ForKeyword(service.Status, c))
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		}).Tag("services")
	})
}
