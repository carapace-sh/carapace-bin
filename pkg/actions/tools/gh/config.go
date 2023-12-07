package gh

import (
	"os"
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/traverse"
	"gopkg.in/yaml.v3"
)

type hostConfig map[string]struct {
	User        string `yaml:"user"`
	OauthToken  string `yaml:"oauth_token"`
	GitProtocol string `yaml:"git_protocol"`
	Users       map[string]struct {
		OauthToken string `yaml:"oauth_token"`
	} `yaml:"users"`
}

func configDir(c carapace.Context) string {
	if v := c.Getenv("GH_CONFIG_DIR"); v != "" {
		return v
	}
	homeDir, _ := traverse.UserHomeDir(c) // TODO handle error
	return filepath.Join(homeDir, ".config", "gh")
}

func actionHostConfig(f func(config hostConfig) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := os.ReadFile(configDir(c) + "/hosts.yml")
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		var config hostConfig
		if err := yaml.Unmarshal(content, &config); err != nil {
			return carapace.ActionMessage(err.Error())
		}
		return f(config)
	})
}

// ActionConfigHosts completes configured hosts
//
//	github.com
//	another.com
func ActionConfigHosts() carapace.Action {
	return actionHostConfig(func(config hostConfig) carapace.Action {
		vals := make([]string, 0)
		for host := range config {
			vals = append(vals, host)
		}
		return carapace.ActionValues(vals...)
	})
}

// ActionConfigUsers completes configured users
//
//	userA
//	userB
func ActionConfigUsers(host string) carapace.Action {
	return actionHostConfig(func(config hostConfig) carapace.Action {
		vals := make([]string, 0)
		for confHost, conf := range config {
			switch host {
			case "", confHost:
				for user := range conf.Users {
					vals = append(vals, user)
				}
			}
		}
		return carapace.ActionValues(vals...)
	})
}
