package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/vercel_completer/cmd/action"
	"github.com/spf13/cobra"
)

var promoteCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote an existing Deployment to current",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promoteCmd).Standalone()

	promoteCmd.Flags().String("timeout", "", "Time to wait for promotion completion")
	promoteCmd.Flags().Bool("yes", false, "Skip confirmation")

	rootCmd.AddCommand(promoteCmd)

	carapace.Gen(promoteCmd).PositionalCompletion(
		action.ActionDeployments(promoteCmd),
	)
}
