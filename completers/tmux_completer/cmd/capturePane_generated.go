package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var capturePaneCmd = &cobra.Command{
	Use:   "capture-pane",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(capturePaneCmd).Standalone()

	capturePaneCmd.Flags().BoolS("C", "C", false, "TODO description")
	capturePaneCmd.Flags().StringS("E", "E", "", "end-line")
	capturePaneCmd.Flags().BoolS("J", "J", false, "TODO description")
	capturePaneCmd.Flags().BoolS("N", "N", false, "TODO description")
	capturePaneCmd.Flags().BoolS("P", "P", false, "TODO description")
	capturePaneCmd.Flags().StringS("S", "S", "", "start-line")
	capturePaneCmd.Flags().BoolS("a", "a", false, "TODO description")
	capturePaneCmd.Flags().StringS("b", "b", "", "buffer-name")
	capturePaneCmd.Flags().BoolS("e", "e", false, "TODO description")
	capturePaneCmd.Flags().BoolS("p", "p", false, "TODO description")
	capturePaneCmd.Flags().BoolS("q", "q", false, "TODO description")
	capturePaneCmd.Flags().StringS("t", "t", "", "target-pane")
	rootCmd.AddCommand(capturePaneCmd)
}
