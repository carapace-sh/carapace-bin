package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:     "undo",
	Short:   "Undo the last operation by reverting to the previous snapshot",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "operation history",
}

func init() {
	carapace.Gen(undoCmd).Standalone()

	undoCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(undoCmd)
}
