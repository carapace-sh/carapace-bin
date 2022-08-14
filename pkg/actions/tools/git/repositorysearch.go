package git

import (
	"strings"

	"github.com/rsteube/carapace"
	gh "github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	glab "github.com/rsteube/carapace-bin/completers/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

// ActionRepositorySearch completes repositories from github.com and gitlab.com
//
//	https://github.com/spf13/cobra
//	https://gitlab.com/gitlab-org/gitlab-runner
func ActionRepositorySearch() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		prefix := "https://"
		if !strings.HasPrefix(c.CallbackValue, prefix) {
			return carapace.ActionValues(prefix).NoSpace()
		}
		c.CallbackValue = strings.TrimPrefix(c.CallbackValue, prefix)

		if !strings.Contains(c.CallbackValue, "/") {
			return carapace.ActionValues("github.com", "gitlab.com").Invoke(c).Prefix(prefix).Suffix("/").ToA().NoSpace()
		}

		// TODO create repo completer with confighost prefix then use carapace.Batch for github.com/gitlab.com
		if strings.HasPrefix(c.CallbackValue, "github.com/") {
			splitted := strings.Split(c.CallbackValue, "/")
			switch len(splitted) {
			case 0:
				return gh.ActionConfigHosts().Invoke(c).Suffix("/").ToA()
			default:
				c.CallbackValue = strings.TrimPrefix(c.CallbackValue, splitted[0]+"/")
				return gh.ActionOwnerRepositories(&cobra.Command{}).Invoke(c).Prefix(prefix + splitted[0] + "/").ToA()
			}
		}

		if strings.HasPrefix(c.CallbackValue, "gitlab.com/") {
			return glab.ActionRepo(&cobra.Command{}).Invoke(c).Prefix(prefix).ToA() // TODO supports hosts
		}
		return carapace.ActionValues()
	})
}
