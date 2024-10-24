package goreleaser

import (
	"os"

	"github.com/carapace-sh/carapace"
	"gopkg.in/yaml.v3"
)

type goreleaserConfig struct {
	Builds []struct {
		Id string
	}
}

// ActionBuilds completes build ids.
//
//	default
//	termux
func ActionBuilds(path string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if path == "" {
			path = ".goreleaser.yaml"
			if _, err := os.Stat(path); err != nil && os.IsNotExist(err) {
				path = ".goreleaser.yml"
			}
		}

		content, err := os.ReadFile(path)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var config goreleaserConfig
		if err := yaml.Unmarshal(content, &config); err != nil {
			return carapace.ActionMessage(err.Error())
		}

		vals := make([]string, 0)
		for _, build := range config.Builds {
			vals = append(vals, build.Id)
		}
		return carapace.ActionValues(vals...)
	}).Tag("builds")
}
