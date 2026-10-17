package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
)

// ActionHosts completes hosts by label
//
//	studio
func ActionHosts() carapace.Action {
	return carapace.ActionExecCommand("rex", "hosts", "--json")(func(output []byte) carapace.Action {
		var hosts struct {
			Home struct {
				Label    string `json:"label"`
				Endpoint string `json:"endpoint"`
			} `json:"home"`
			Hosts []struct {
				Label    string `json:"label"`
				Endpoint string `json:"endpoint"`
			} `json:"hosts"`
		}
		if err := json.Unmarshal(output, &hosts); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, host := range hosts.Hosts {
			if host.Label != "" {
				vals = append(vals, host.Label, host.Endpoint)
			}
		}
		if hosts.Home.Label != "" {
			vals = append(vals, hosts.Home.Label, hosts.Home.Endpoint)
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(24 * time.Hour).Tag("hosts")
}
