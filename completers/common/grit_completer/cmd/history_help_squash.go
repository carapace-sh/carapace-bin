package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_squashCmd = &cobra.Command{
	Use:   "squash",
	Short: "Fold a revision range into the oldest commit and replay descendants",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_squashCmd).Standalone()

	history_helpCmd.AddCommand(history_help_squashCmd)
}
