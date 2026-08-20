package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_saveSessionCmd = &cobra.Command{
	Use:   "save-session",
	Short: "Save the current session state to disk immediately",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_saveSessionCmd).Standalone()

	action_saveSessionCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_saveSessionCmd)
}
