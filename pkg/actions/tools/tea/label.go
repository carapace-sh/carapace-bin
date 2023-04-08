package tea

import (
	"os/exec"
	"strings"

	"github.com/rsteube/carapace"
	"gopkg.in/yaml.v3"
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
		args := []string{"label", "list", "--output", "yaml", "--limit", "100"}
		if opts.Login != "" {
			args = append(args, "--login", opts.Login)
		}
		if opts.Remote != "" {
			args = append(args, "--remote", opts.Remote)
		}
		if opts.Repo != "" {
			args = append(args, "--repo", opts.Repo)
		}

		return carapace.ActionExecCommandE("tea", args...)(func(output []byte, err error) carapace.Action {
			if splitted := strings.SplitN(string(output), "\n", 2); strings.HasPrefix(splitted[0], "NOTE:") {
				switch len(splitted) {
				case 1:
					output = []byte{}
				default:
					output = []byte(splitted[1])
				}
			}

			if err != nil {
				if _, ok := err.(*exec.ExitError); ok {
					return carapace.ActionMessage(strings.SplitN(string(output), "\n", 2)[0])
				}
				return carapace.ActionMessage(err.Error())
			}

			var labels []struct {
				Name        string `yaml:"Name"`
				Description string `yaml:"Description"`
				Color       string `yaml:"Color"`
			}

			if err := yaml.Unmarshal(output, &labels); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for _, label := range labels {
				vals = append(vals, label.Name, label.Description, "#"+label.Color)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
