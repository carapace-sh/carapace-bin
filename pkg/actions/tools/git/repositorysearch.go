package git

import (
	"strings"

	"github.com/carapace-sh/carapace"
	gh "github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"     // TODO move to pkg/actions/tools
	glab "github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action" // TODO move to pkg/actions/tools
	"github.com/spf13/cobra"
)

type SearchOpts struct {
	Prefix bool
	Github bool
	Gitlab bool
}

func (o SearchOpts) Default() SearchOpts {
	o.Prefix = true
	o.Github = true
	o.Gitlab = true
	return o
}

func (o SearchOpts) Hosts() []string {
	hosts := make([]string, 0)
	if o.Github {
		hosts = append(hosts, "github.com")
	}
	if o.Gitlab {
		hosts = append(hosts, "gitlab.com")
	}
	return hosts
}

// ActionRepositorySearch completes repositories from github.com and gitlab.com
//
//	https://github.com/spf13/cobra
//	https://gitlab.com/gitlab-org/gitlab-runner
func ActionRepositorySearch(opts SearchOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		prefix := ""
		if opts.Prefix {
			prefix = "https://"
			if !strings.HasPrefix(c.Value, prefix) {
				return carapace.ActionValues(prefix).NoSpace()
			}
			c.Value = strings.TrimPrefix(c.Value, prefix)
		}

		if !strings.Contains(c.Value, "/") {
			return carapace.ActionValues(opts.Hosts()...).Invoke(c).Prefix(prefix).Suffix("/").ToA().NoSpace()
		}

		// TODO create repo completer with confighost prefix then use carapace.Batch for github.com/gitlab.com
		switch strings.SplitAfter(c.Value, "/")[0] {
		case "github.com/":
			splitted := strings.Split(c.Value, "/")
			switch len(splitted) {
			case 0:
				return gh.ActionConfigHosts().Invoke(c).Suffix("/").ToA()
			default:
				c.Value = strings.TrimPrefix(c.Value, splitted[0]+"/")
				return gh.ActionOwnerRepositories(&cobra.Command{}).Invoke(c).Prefix(prefix + splitted[0] + "/").ToA()
			}

		case "gitlab.com/":
			return glab.ActionRepo(&cobra.Command{}).Invoke(c).Prefix(prefix).ToA() // TODO supports hosts

		default:
			return carapace.ActionValues()
		}
	})
}
