package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_listPanesCmd = &cobra.Command{
	Use:   "list-panes",
	Short: "List all panes in the current session",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_listPanesCmd).Standalone()

	action_listPanesCmd.Flags().BoolP("all", "a", false, "Include all available fields")
	action_listPanesCmd.Flags().BoolP("command", "c", false, "Include running command information")
	action_listPanesCmd.Flags().BoolP("geometry", "g", false, "Include geometry (position, size)")
	action_listPanesCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_listPanesCmd.Flags().BoolP("json", "j", false, "Output as JSON")
	action_listPanesCmd.Flags().BoolP("state", "s", false, "Include pane state (focused, floating, exited, etc.)")
	action_listPanesCmd.Flags().BoolP("tab", "t", false, "Include tab information (name, position, ID)")
	actionCmd.AddCommand(action_listPanesCmd)
}
