package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var ruleset_viewCmd = &cobra.Command{
	Use:   "view [<ruleset-id>]",
	Short: "View information about a ruleset",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ruleset_viewCmd).Standalone()

	ruleset_viewCmd.Flags().StringP("org", "o", "", "Organization name if the provided ID is an organization-level ruleset")
	ruleset_viewCmd.Flags().BoolP("parents", "p", false, "Whether to include rulesets configured at higher levels that also apply")
	ruleset_viewCmd.Flags().BoolP("web", "w", false, "Open the ruleset in the browser")
	rulesetCmd.AddCommand(ruleset_viewCmd)

	carapace.Gen(ruleset_viewCmd).FlagCompletion(carapace.ActionMap{
		"org": gh.ActionOrganizations(gh.HostOpts{}),
	})

	// TODO organization rulesets
	carapace.Gen(ruleset_viewCmd).PositionalCompletion(
		action.ActionRulesets(ruleset_viewCmd),
	)
}
