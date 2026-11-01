package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_help_renameTabByIdCmd = &cobra.Command{
	Use:   "rename-tab-by-id",
	Short: "Rename tab by stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_help_renameTabByIdCmd).Standalone()

	action_helpCmd.AddCommand(action_help_renameTabByIdCmd)
}
