package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_listTabsCmd = &cobra.Command{
	Use:   "list-tabs",
	Short: "List all tabs with their information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_listTabsCmd).Standalone()

	action_listTabsCmd.Flags().BoolP("all", "a", false, "Include all available fields")
	action_listTabsCmd.Flags().BoolP("dimensions", "d", false, "Include dimension information (viewport, display area)")
	action_listTabsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	action_listTabsCmd.Flags().BoolP("json", "j", false, "Output as JSON")
	action_listTabsCmd.Flags().BoolP("layout", "l", false, "Include layout information (swap layout name and dirty state)")
	action_listTabsCmd.Flags().BoolP("panes", "p", false, "Include pane counts")
	action_listTabsCmd.Flags().BoolP("state", "s", false, "Include state information (active, fullscreen, sync, floating visibility)")
	actionCmd.AddCommand(action_listTabsCmd)
}
