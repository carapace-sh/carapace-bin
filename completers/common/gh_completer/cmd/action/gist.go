package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// gistApiHost strips the `gist.` prefix from a gist url host,
// returning the corresponding api host.
func gistApiHost(host string) string {
	return strings.TrimPrefix(host, "gist.")
}

func ActionGistUrls(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		prefix := "https://gist."
		if !strings.HasPrefix(c.Value, prefix) {
			// ActionMultiParts to force nospace (should maybe added as function to InvokedAction to prevent space suffix)
			return carapace.ActionMultiParts(".", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return carapace.ActionValues(prefix)
				default:
					return carapace.ActionValues()
				}
			})
		} else {
			c.Value = strings.TrimPrefix(c.Value, prefix)
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				switch len(c.Parts) {
				case 0:
					return ActionConfigHosts().Invoke(c).Suffix("/").ToA()
				case 1:
					return gh.ActionUsers(gh.HostOpts{}).Invoke(c).Suffix("/").ToA()
				case 2:
					return gh.ActionGistIds(gh.GistOpts{Host: gistApiHost(c.Parts[0]), Owner: c.Parts[1]})
				default:
					return carapace.ActionValues()
				}
			}).Invoke(c).Prefix(prefix).ToA()
		}
	})
}

func ActionGistIds(cmd *cobra.Command) carapace.Action {
	return gh.ActionGistIds(gh.GistOpts{})
}

func ActionGists(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.HasPrefix(c.Value, "h") {
			return ActionGistUrls(cmd)
		} else {
			return ActionGistIds(cmd)
		}
	})
}

func ActionGistFiles(cmd *cobra.Command, name string) carapace.Action {
	splitted := strings.Split(name, "/")
	gistId := splitted[len(splitted)-1]
	var host string
	var owner string
	switch len(splitted) {
	case 1:
		host = gistApiHost(splitted[2])
		owner = splitted[3]
	default:
		return carapace.ActionMessage("invalid gist id/url: " + name)
	}
	return gh.ActionGistFiles(gh.GistOpts{Host: host, Owner: owner, Gist: gistId})
}
