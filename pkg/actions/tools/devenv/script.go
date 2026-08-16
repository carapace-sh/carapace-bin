package devenv

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionScripts completes scripts defined in devenv.nix
//
//	hello (greet the user)
//	migrate (run pending migrations)
func ActionScripts(opts GlobalOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		file, err := configFile(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionExecCommand("devenv", opts.command("eval", "scripts")...)(func(output []byte) carapace.Action {
			var evaluated struct {
				Scripts map[string]struct {
					Description string `json:"description"`
				} `json:"scripts"`
			}
			if err := json.Unmarshal(output, &evaluated); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for name, script := range evaluated.Scripts {
				vals = append(vals, name, script.Description)
			}
			return carapace.ActionValuesDescribed(vals...)
		}).Cache(5*time.Minute, opts.cacheKey(), key.FileStats(file))
	}).Tag("scripts")
}
