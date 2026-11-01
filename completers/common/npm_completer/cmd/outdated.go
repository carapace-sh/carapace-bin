package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var outdatedCmd = &cobra.Command{
	Use:   "outdated",
	Short: "Check for outdated packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outdatedCmd).Standalone()
	outdatedCmd.Flags().BoolP("all", "a", false, "show all")
	outdatedCmd.Flags().String("before", "", "only show versions available on or before the given date")
	outdatedCmd.Flags().BoolP("global", "g", false, "operate globally")
	outdatedCmd.Flags().Bool("json", false, "output as json")
	outdatedCmd.Flags().BoolP("long", "l", false, "show extended information")
	outdatedCmd.Flags().String("min-release-age", "", "only show versions available more than the given number of days ago")
	outdatedCmd.Flags().StringArray("min-release-age-exclude", nil, "packages exempt from the min-release-age filter")
	outdatedCmd.Flags().BoolP("parseable", "p", false, "output parseable results")
	addWorkspaceFlags(outdatedCmd)

	rootCmd.AddCommand(outdatedCmd)

	carapace.Gen(outdatedCmd).PositionalAnyCompletion(
		action.ActionPackages(outdatedCmd),
	)
}
