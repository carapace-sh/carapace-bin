package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_action_renameTabByIdCmd = &cobra.Command{
	Use:   "rename-tab-by-id",
	Short: "Rename tab by stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_action_renameTabByIdCmd).Standalone()

	help_actionCmd.AddCommand(help_action_renameTabByIdCmd)
}
