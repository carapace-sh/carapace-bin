package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_renameTabByIdCmd = &cobra.Command{
	Use:   "rename-tab-by-id",
	Short: "Rename tab by stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_renameTabByIdCmd).Standalone()

	action_renameTabByIdCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_renameTabByIdCmd)
}
