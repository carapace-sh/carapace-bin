package ssh

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-ssh/config"
)

// ActionHosts completes ssh hosts.
func ActionHosts(configPath string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		opts := make([]config.Option, 0, 1)
		if home := homeDir(c); home != "" {
			opts = append(opts, config.WithHomeDir(home))
		}

		hosts, _ := config.HostCandidates(configPath, opts...)
		return carapace.ActionValues(hosts...).Tag("ssh hosts")
	})
}

func homeDir(c carapace.Context) string {
	if h, ok := c.LookupEnv("HOME"); ok && h != "" {
		return h
	}
	if h, err := os.UserHomeDir(); err == nil {
		return h
	}
	return ""
}
