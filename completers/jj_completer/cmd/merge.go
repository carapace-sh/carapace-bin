package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "Merge work from multiple branches",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mergeCmd).Standalone()

	mergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	mergeCmd.Flags().BoolP("insert-after", "A", false, "Insert the new change between the target commit(s) and their children")
	mergeCmd.Flags().BoolP("insert-before", "B", false, "Insert the new change between the target commit(s) and their parents")
	mergeCmd.Flags().StringSliceP("message", "m", []string{}, "The change description to use")
	rootCmd.AddCommand(mergeCmd)
}
