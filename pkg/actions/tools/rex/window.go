package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// WindowsOpts provides scope filters for window completion
type WindowsOpts struct {
	Server  string // host label or server address
	Session string // session target
}

func (o WindowsOpts) cacheKey() key.Key {
	return key.String(o.Server, o.Session)
}

// ActionWindows completes window targets
//
//	window:01a14a6c-c12e-7af4-be34-daeabea276d6 (Window bea276d6)
func ActionWindows(opts WindowsOpts) carapace.Action {
	args := []string{"window", "ls", "--json"}
	if opts.Session != "" {
		args = append(args, "--session", opts.Session)
	}
	if opts.Server != "" {
		args = append(args, "--server", opts.Server)
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
	}).Cache(time.Minute, opts.cacheKey()).Tag("windows")
}
