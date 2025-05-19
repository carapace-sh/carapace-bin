package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sparse_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the patterns that are currently present in the working copy",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sparse_listCmd).Standalone()

	sparse_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	sparseCmd.AddCommand(sparse_listCmd)
}
