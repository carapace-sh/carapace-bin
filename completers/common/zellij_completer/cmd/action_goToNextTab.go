package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_goToNextTabCmd = &cobra.Command{
	Use:   "go-to-next-tab",
	Short: "Go to the next tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_goToNextTabCmd).Standalone()

	action_goToNextTabCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_goToNextTabCmd)
}
