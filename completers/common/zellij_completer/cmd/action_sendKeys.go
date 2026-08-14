package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_sendKeysCmd = &cobra.Command{
	Use:   "send-keys",
	Short: "Send one or more keys to the terminal (e.g., \"Ctrl a\", \"F1\", \"Alt Shift b\")",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_sendKeysCmd).Standalone()

	action_sendKeysCmd.Flags().BoolP("help", "h", false, "Print help")
	action_sendKeysCmd.Flags().StringP("pane-id", "p", "", "The pane_id of the pane, eg. terminal_1, plugin_2 or 3 (equivalent to terminal_3)")
	actionCmd.AddCommand(action_sendKeysCmd)
}
