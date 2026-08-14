package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_renameSessionCmd = &cobra.Command{
	Use:   "rename-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_renameSessionCmd).Standalone()

	action_renameSessionCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_renameSessionCmd)
}
