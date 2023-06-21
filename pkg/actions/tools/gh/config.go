package gh

import (
	"os"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/gh/config"
	"github.com/rsteube/carapace/pkg/xdg"
	"gopkg.in/yaml.v3"
)

func userFor(host string) (string, error) {
	if host == "" {
		host = "github.com"
	}

	configDir, err := xdg.UserConfigDir()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(configDir + "/gh/hosts.yml")
	if err != nil {
		return "", err
	}

	var hosts map[string]struct {
		User string
	}
	if err := yaml.Unmarshal(content, &hosts); err != nil {
		return "", err
	}

	return hosts[host].User, nil
}

// ActionConfigHosts completes configured hosts
//
//	github.com
//	another.com
func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		} else {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage("failed ot loadd hosts: " + err.Error())
			} else {
				return carapace.ActionValues(hosts...)
			}
		}
	})
}
