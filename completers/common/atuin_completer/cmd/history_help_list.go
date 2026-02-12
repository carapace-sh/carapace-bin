package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var history_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items in history",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(history_help_listCmd).Standalone()

	history_helpCmd.AddCommand(history_help_listCmd)
}
