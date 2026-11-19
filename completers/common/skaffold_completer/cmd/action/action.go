package action

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
	"go.yaml.in/yaml/v3"
)

type config struct {
	Profiles []struct {
		Name string `yaml:"name"`
	} `yaml:"profiles"`
}

// ActionProfiles completes profiles of the skaffold config
func ActionProfiles(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		configFile := "skaffold.yaml"
		if flag := cmd.Flag("filename"); flag != nil && flag.Value.String() != "" {
			configFile = flag.Value.String()
		}

		content, err := os.ReadFile(configFile)
		if err != nil {
			return carapace.ActionValues()
		}

		var cfg config
		if err := yaml.Unmarshal(content, &cfg); err != nil {
			return carapace.ActionValues()
		}

		vals := make([]string, 0)
		for _, profile := range cfg.Profiles {
			vals = append(vals, profile.Name)
		}
		return carapace.ActionValues(vals...).Unique()
	}).Tag("profiles")
}
