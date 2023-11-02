package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var restoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore paths from another revision",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(restoreCmd).Standalone()

	restoreCmd.Flags().StringP("changes-in", "c", "", "Undo the changes in a revision as compared to the merge of its parents")
	restoreCmd.Flags().String("from", "", "Revision to restore from (source)")
	restoreCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	restoreCmd.Flags().String("to", "", "Revision to restore into (destination)")
	rootCmd.AddCommand(restoreCmd)
}
