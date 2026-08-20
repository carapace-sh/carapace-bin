package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_goToTabCmd = &cobra.Command{
	Use:   "go-to-tab",
	Short: "Go to tab with index [index]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_goToTabCmd).Standalone()

	action_goToTabCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_goToTabCmd)
}
