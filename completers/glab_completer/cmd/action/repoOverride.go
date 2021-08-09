package action

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
	"strings"
)

func ActionRepo(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		configHosts, err := hosts()
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}

		if strings.Contains(c.CallbackValue, "/") {
			isKnownHost := false
			for _, host := range configHosts {
				if strings.HasPrefix(c.CallbackValue, host) {
					isKnownHost = true
					break
				}
			}
			if !isKnownHost {
				return carapace.ActionValues() // only complete full host style for simplicity
			}
		}

		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return carapace.ActionValues(configHosts...).Invoke(c).Suffix("/").ToA()
			case 1:
				return carapace.Batch(
					ActionUsers(cmd),
					ActionGroups(cmd),
				).Invoke(c).Merge().Suffix("/").ToA()
			case 2:
				b := carapace.Batch(
					ActionSubgroups(cmd, c.Parts[1]),
					ActionGroupProjects(cmd, c.Parts[1]),
					ActionUserProjects(cmd, c.Parts[1]),
				).Invoke(c)
				return b[0].Suffix("/").Merge(b[1:]...).ToA() // ActionSubgroups needs `/` suffix
			case 3:
				groupProjects := ActionGroupProjects(cmd, strings.Join(c.Parts[1:], "/")).Invoke(c)
				return groupProjects.ToA()
			default:
				return carapace.ActionValues()
			}
		})
	})
}
