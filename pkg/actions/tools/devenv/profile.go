package devenv

import (
	"encoding/json"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

// ActionProfiles completes profiles defined in devenv.nix.
// The `hostname` and `user` attributes are skipped as these only contain automatically activated profiles.
//
//	backend
//	frontend
func ActionProfiles(opts GlobalOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		file, err := configFile(c)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		return carapace.ActionExecCommand("devenv", opts.command("eval", "profiles")...)(func(output []byte) carapace.Action {
			var evaluated struct {
				Profiles map[string]json.RawMessage `json:"profiles"`
			}
			if err := json.Unmarshal(output, &evaluated); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for name := range evaluated.Profiles {
				switch name {
				case "hostname", "user":
				default:
					vals = append(vals, name)
				}
			}
			return carapace.ActionValues(vals...)
		}).Cache(5*time.Minute, opts.cacheKey(), key.FileStats(file))
	}).Tag("profiles")
}
