package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var breakPaneCmd = &cobra.Command{
	Use:   "break-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(breakPaneCmd).Standalone()

	breakPaneCmd.Flags().StringS("F", "F", "", "format")
	breakPaneCmd.Flags().BoolS("P", "P", false, "TODO description")
	breakPaneCmd.Flags().BoolS("a", "a", false, "TODO description")
	breakPaneCmd.Flags().BoolS("b", "b", false, "TODO description")
	breakPaneCmd.Flags().BoolS("d", "d", false, "TODO description")
	breakPaneCmd.Flags().StringS("n", "n", "", "window-name")
	breakPaneCmd.Flags().StringS("s", "s", "", "src-pane")
	breakPaneCmd.Flags().StringS("t", "t", "", "dst-window")
	rootCmd.AddCommand(breakPaneCmd)
}
