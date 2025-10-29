package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var target_help_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed and available targets",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(target_help_listCmd).Standalone()

	target_helpCmd.AddCommand(target_help_listCmd)
}
