package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var ruleset_checkCmd = &cobra.Command{
	Use:   "check [<branch>]",
	Short: "View rules that would apply to a given branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ruleset_checkCmd).Standalone()

	ruleset_checkCmd.Flags().Bool("default", false, "Check rules on default branch")
	ruleset_checkCmd.Flags().BoolP("web", "w", false, "Open the branch rules page in a web browser")
	rulesetCmd.AddCommand(ruleset_checkCmd)

	carapace.Gen(ruleset_checkCmd).PositionalCompletion(
		action.ActionBranches(ruleset_checkCmd),
	)
}
