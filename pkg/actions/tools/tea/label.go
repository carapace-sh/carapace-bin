package tea

import (
	"github.com/rsteube/carapace"
)

type RepoOpts struct {
	Login  string
	Remote string
	Repo   string
}

// ActionLabels completes labels
//
//	Cleanup (Cleanup and cosmetic)
//	Discussion
func ActionLabels(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var labels []struct {
			Name        string `yaml:"Name"`
			Description string `yaml:"Description"`
			Color       string `yaml:"Color"`
		}

		return actionYamlQuery(opts, &labels, "label", "list")(func() carapace.Action {
			vals := make([]string, 0)
			for _, label := range labels {
				vals = append(vals, label.Name, label.Description, "#"+label.Color)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
