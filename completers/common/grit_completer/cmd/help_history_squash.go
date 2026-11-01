package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_history_squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Fold a revision range into the oldest commit and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_history_squashCmd).Standalone()

	help_historyCmd.AddCommand(help_history_squashCmd)
}
