package action

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

func ActionBodyLinks(cmd *cobra.Command) carapace.Action {
	return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
		keywords := []string{"close", "closes", "closed", "fix", "fixes", "fixed", "resolve", "resolves", "resolved"} // official keywords
		keywords = append(keywords, "causes", "caused-by", "related")                                                 // additional keywords

		re := regexp.MustCompile(fmt.Sprintf("^(%v)$", strings.Join(keywords, "|")))
		if last := len(c.Parts) - 1; last < 0 || !re.MatchString(c.Parts[last]) {
			return carapace.ActionValues(keywords...).Invoke(c).Suffix(" #").ToA()
		}

		return carapace.ActionMultiParts("#", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				if strings.Contains(c.CallbackValue, "/") {
					return ActionOwnerRepositories(cmd).Invoke(c).Suffix("#").ToA()
				}
				return ActionOwnerRepositories(cmd)
			case 1:
				fakeCmd := cobra.Command{}
				fakeCmd.Flags().String("repo", c.Parts[0], "")
				fakeCmd.Flag("repo").Changed = true
				return carapace.Batch(
					ActionIssues(&fakeCmd, IssueOpts{Open: true, Closed: true}),
					ActionPullRequests(&fakeCmd, PullRequestOpts{Open: true, Closed: true, Merged: true}),
				).ToA()
			default:
				return carapace.ActionValues()
			}
		})
	})
}
