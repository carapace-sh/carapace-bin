package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var ruleset_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List rulesets for a repository or organization",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ruleset_listCmd).Standalone()

	ruleset_listCmd.Flags().StringP("limit", "L", "", "Maximum number of rulesets to list")
	ruleset_listCmd.Flags().StringP("org", "o", "", "List organization-wide rulesets for the provided organization")
	ruleset_listCmd.Flags().BoolP("parents", "p", false, "Whether to include rulesets configured at higher levels that also apply")
	ruleset_listCmd.Flags().BoolP("web", "w", false, "Open the list of rulesets in the web browser")
	rulesetCmd.AddCommand(ruleset_listCmd)

	carapace.Gen(ruleset_listCmd).FlagCompletion(carapace.ActionMap{
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})
}
