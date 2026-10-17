package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionWindows completes window targets
//
//	window:01a14a6c-c12e-7af4-be34-daeabea276d6 (Window bea276d6)
func ActionWindows(server, session string) carapace.Action {
	args := []string{"window", "ls", "--json"}
	if session != "" {
		args = append(args, "--session", session)
	}
	if server != "" {
		args = append(args, "--server", server)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		var windows struct {
			Windows []struct {
				WindowId string `json:"window_id"`
				Label    string `json:"label"`
			}
		}
		if err := json.Unmarshal(output, &windows); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0, len(windows.Windows)*2)
		for _, window := range windows.Windows {
			if window.Label != "" {
				vals = append(vals, window.WindowId, window.Label)
			} else {
				vals = append(vals, window.WindowId)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(time.Minute, key.String(server, session)).Tag("windows")
}
