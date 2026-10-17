package rex

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionBlockFlavors completes block flavors for `rex block new`
//
//	com.superlogical.terminal.shell (Shell)
func ActionBlockFlavors(server string) carapace.Action {
	args := []string{"block", "types", "--json"}
	if server != "" {
		args = append(args, "--server", server)
	}
	return carapace.ActionExecCommand("rex", args...)(func(output []byte) carapace.Action {
		var types struct {
			Creators []struct {
				Name    string `json:"name"`
				Flavors []struct {
					Name  string `json:"name"`
					Title string `json:"title"`
				} `json:"flavors"`
			} `json:"creators"`
		}
		if err := json.Unmarshal(output, &types); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		vals := make([]string, 0)
		for _, creator := range types.Creators {
			for _, flavor := range creator.Flavors {
				vals = append(vals, creator.Name+"."+flavor.Name, flavor.Title)
			}
		}
		return carapace.ActionValuesDescribed(vals...)
	}).Cache(24*time.Hour, key.String(server)).Tag("flavors")
}
