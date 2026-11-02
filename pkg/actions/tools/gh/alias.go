package gh

import (
	"errors"
	"maps"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh/config"
)

func Aliases() (map[string]string, error) {
	if config, err := config.ParseDefaultConfig(); err != nil {
		return nil, errors.New("failed to parse DefaultConfig: " + err.Error())
	} else {
		if aliasCfg, err := config.Aliases(); err != nil {
			return nil, errors.New("failed to load AliasCfg: " + err.Error())
		} else {
			aliases := make(map[string]string)
			maps.Copy(aliases, aliasCfg.All())
			return aliases, nil
		}
	}
}

// ActionAliases completes gh aliases
//
//	co (pr checkout)
func ActionAliases() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if aliases, err := Aliases(); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			values := make([]string, 0)
			for key, value := range aliases {
				values = append(values, key, value)
			}
			return carapace.ActionValuesDescribed(values...)
		}
	}).Tag("aliases")
}
