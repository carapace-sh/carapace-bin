package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_overrideLayoutCmd = &cobra.Command{
	Use:   "override-layout",
	Short: "Override the layout of the active tab",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_overrideLayoutCmd).Standalone()

	action_overrideLayoutCmd.Flags().Bool("apply-only-to-active-tab", false, "Only apply the layout to the active tab (uses just the first layout tab if it has multiple)")
	action_overrideLayoutCmd.Flags().BoolP("help", "h", false, "Print help")
	action_overrideLayoutCmd.Flags().String("layout-dir", "", "Default folder to look for layouts")
	action_overrideLayoutCmd.Flags().String("layout-string", "", "Raw KDL layout string to use directly (instead of a layout file path)")
	action_overrideLayoutCmd.Flags().Bool("retain-existing-plugin-panes", false, "Retain existing plugin panes that do not fit with the layout default: false)")
	action_overrideLayoutCmd.Flags().Bool("retain-existing-terminal-panes", false, "Retain existing terminal panes that do not fit in the layout (default: false)")
	actionCmd.AddCommand(action_overrideLayoutCmd)

	carapace.Gen(action_overrideLayoutCmd).FlagCompletion(carapace.ActionMap{
		"layout-dir": carapace.ActionFiles(),
	})

	carapace.Gen(action_overrideLayoutCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
