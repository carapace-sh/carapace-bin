package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_goToTabNameCmd = &cobra.Command{
	Use:   "go-to-tab-name",
	Short: "Go to tab with name [name]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_goToTabNameCmd).Standalone()

	action_goToTabNameCmd.Flags().BoolP("create", "c", false, "Create a tab if one does not exist")
	action_goToTabNameCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	actionCmd.AddCommand(action_goToTabNameCmd)
}
