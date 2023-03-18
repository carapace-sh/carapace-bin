package yarn

import (
	"encoding/json"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

// ActionConfigKeys completes configuration keys
//
//	npmPublishRegistry (Registry to push packages to)
//	npmRegistries (Settings per registry)
func ActionConfigKeys() carapace.Action {
	return carapace.ActionExecCommand("yarn", "config", "--json")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		vals := make([]string, 0)

		for _, line := range lines[:len(lines)-1] {
			var config struct {
				Key         string
				Description string
			}
			if err := json.Unmarshal([]byte(line), &config); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			vals = append(vals, config.Key, config.Description)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}

// ActionConfigValues completes configuration options for given key
//
//	true
//	strict
func ActionConfigValues(key string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO add missing config options
		switch key {
		case "pnpMode":
			return carapace.ActionValuesDescribed(
				"strict", "generate standard PnP maps",
				"loose", "merge with the n_m resolution",
			).StyleF(style.ForKeyword)
		case "checksumBehavior":
			return carapace.ActionStyledValuesDescribed(
				"throw", "throw an exception", style.Carapace.KeywordNegative,
				"update", "update the lockfile checksum", style.Carapace.KeywordPositive,
				"ignore", "the checksum check will not happen", style.Carapace.KeywordAmbiguous,
			)

		case "nodeLinker":
			return carapace.ActionValues("pnp", "pnpm", "node-modules")

		default:
			return carapace.ActionExecCommand("yarn", "config", "--json")(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")

				for _, line := range lines[:len(lines)-1] {
					var config struct {
						Key  string
						Type string
					}
					if err := json.Unmarshal([]byte(line), &config); err != nil {
						return carapace.ActionMessage(err.Error())
					}

					if config.Key == key {
						switch config.Type {
						case "ABSOLUTE_PATH":
							return carapace.ActionFiles()
						case "BOOLEAN":
							return carapace.ActionValues("true", "false").StyleF(style.ForKeyword)
						default:
							return carapace.ActionValues()
						}
					}
				}
				return carapace.ActionValues()
			})
		}
	})
}
