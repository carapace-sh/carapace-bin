package ssh

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-ssh/config"
	"github.com/carapace-sh/carapace-ssh/knownhosts"
)

// ActionHosts completes ssh hosts.
func ActionHosts(configPath string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		home := homeDir(c)

		configOpts := make([]config.Option, 0, 1)
		if home != "" {
			configOpts = append(configOpts, config.WithHomeDir(home))
		}

		idx, _ := config.IndexConfig(configPath, configOpts...)
		hosts := idx.HostCandidates()

		knownOpts := make([]knownhosts.Option, 0, 1)
		if home != "" {
			knownOpts = append(knownOpts, knownhosts.WithHomeDir(home))
		}
		knownFromConfig, _ := knownhosts.IndexFromConfigPaths(idx.KnownHostsFiles(), knownOpts...)
		if configPath == "" {
			knownDefault, _ := knownhosts.IndexDefault(knownOpts...)
			hosts = knownhosts.Merge(hosts, knownFromConfig.HostCandidates(), knownDefault.HostCandidates())
		} else if len(knownFromConfig.HostCandidates()) > 0 {
			hosts = knownhosts.Merge(hosts, knownFromConfig.HostCandidates())
		}

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
