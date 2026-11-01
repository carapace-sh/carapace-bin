package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_goToPreviousTabCmd = &cobra.Command{
	Use:   "go-to-previous-tab",
	Short: "Go to the previous tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_goToPreviousTabCmd).Standalone()

	action_goToPreviousTabCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_goToPreviousTabCmd)
}
