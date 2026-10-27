package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_rewordCmd = &cobra.Command{
	Use:   "reword",
	Short: "Replace a commit message and replay descendants onto the new history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_rewordCmd).Standalone()

	history_helpCmd.AddCommand(history_help_rewordCmd)
}
