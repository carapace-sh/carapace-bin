package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_closeTabByIdCmd = &cobra.Command{
	Use:   "close-tab-by-id",
	Short: "Close tab with stable ID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_closeTabByIdCmd).Standalone()

	action_closeTabByIdCmd.Flags().BoolP("help", "h", false, "Print help")
	actionCmd.AddCommand(action_closeTabByIdCmd)
}
