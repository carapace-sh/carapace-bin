package dbt

import (
	"os"

	"github.com/carapace-sh/carapace"
	"gopkg.in/yaml.v3"
)

type profile struct {
	Outputs map[string]struct {
		Type string
	}
}

func actionProfiles(dir string, f func(profiles map[string]profile) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if dir == "" {
			var err error
			if dir, err = os.UserHomeDir(); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			dir += "/.dbt"
		}

		content, err := os.ReadFile(dir + "/profiles.yml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var profiles map[string]profile
		if err := yaml.Unmarshal(content, &profiles); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(profiles)
	})
}

// ActionProfiles completes profiles
//
//	profile1
//	profile2
func ActionProfiles(dir string) carapace.Action {
	return actionProfiles(dir, func(profiles map[string]profile) carapace.Action {
		vals := make([]string, 0)
		for profile := range profiles {
			vals = append(vals, profile)
		}
		return carapace.ActionValues(vals...)
	})
}
