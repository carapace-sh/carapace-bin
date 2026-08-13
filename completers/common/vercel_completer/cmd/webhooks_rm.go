package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var webhooks_rmCmd = &cobra.Command{
	Use:     "rm",
	Aliases: []string{"remove", "delete"},
	Short:   "Remove a webhook",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(webhooks_rmCmd).Standalone()

	webhooks_rmCmd.Flags().Bool("yes", false, "Skip confirmation")

	webhooksCmd.AddCommand(webhooks_rmCmd)

	carapace.Gen(webhooks_rmCmd).PositionalCompletion(
		carapace.ActionValues(),
	)
}
