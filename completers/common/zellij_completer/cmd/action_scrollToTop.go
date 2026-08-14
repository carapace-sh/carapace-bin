package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var action_scrollToTopCmd = &cobra.Command{
	Use:   "scroll-to-top",
	Short: "Scroll up to top in focus pane",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(action_scrollToTopCmd).Standalone()

	action_scrollToTopCmd.Flags().BoolP("help", "h", false, "Print help")
	action_scrollToTopCmd.Flags().StringP("pane-id", "p", "", "Target a specific pane by ID (eg. terminal_1, plugin_2, or 3)")
	actionCmd.AddCommand(action_scrollToTopCmd)
}
