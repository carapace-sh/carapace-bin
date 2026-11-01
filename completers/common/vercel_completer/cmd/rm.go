package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove"},
	Short:   "Remove deployment(s) by project name or deployment ID",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().Bool("hard", false, "Hard delete")
	rmCmd.Flags().BoolP("safe", "s", false, "Skip deployments with an active alias")
	rmCmd.Flags().BoolP("yes", "y", false, "Skip confirmation")

	rootCmd.AddCommand(rmCmd)

	carapace.Gen(rmCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionProjects(rmCmd),
			action.ActionDeployments(rmCmd),
		).ToA(),
	)
}
