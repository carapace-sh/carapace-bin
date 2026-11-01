package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionClients completes client targets
//
//	client:01a14a6c-c089-7cf4-82fb-58f066892623 (qemu's iMac)
func ActionClients(server string) carapace.Action {
	args := []string{"client", "ls", "--json"}
	if server != "" {
		args = append(args, "--server", server)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		var clients struct {
			Clients []struct {
				ClientId string `json:"client_id"`
				Info     struct {
					Label string `json:"label"`
				} `json:"info"`
			}
		}
		if err := json.Unmarshal(output, &clients); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0, len(clients.Clients)*2)
		for _, client := range clients.Clients {
			if client.Info.Label != "" {
				vals = append(vals, client.ClientId, client.Info.Label)
			} else {
				vals = append(vals, client.ClientId)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(time.Minute, key.String(server)).Tag("clients")
}
