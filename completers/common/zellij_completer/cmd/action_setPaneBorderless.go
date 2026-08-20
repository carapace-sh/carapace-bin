package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_setPaneBorderlessCmd = &cobra.Command{
	Use:   "set-pane-borderless",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_setPaneBorderlessCmd).Standalone()

	action_setPaneBorderlessCmd.Flags().BoolP("borderless", "b", false, "Whether the pane should be borderless (flag present) or bordered (flag absent)")
	action_setPaneBorderlessCmd.Flags().BoolP("help", "h", false, "Print help")
	action_setPaneBorderlessCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3)")
	action_setPaneBorderlessCmd.MarkFlagRequired("pane-id")
	actionCmd.AddCommand(action_setPaneBorderlessCmd)
}
