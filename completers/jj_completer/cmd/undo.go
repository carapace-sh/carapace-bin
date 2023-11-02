package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Undo an operation (shortcut for `jj op undo`)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(undoCmd).Standalone()

	undoCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	undoCmd.Flags().StringSlice("what", []string{}, "What portions of the local state to restore (can be repeated)")
	rootCmd.AddCommand(undoCmd)
}
