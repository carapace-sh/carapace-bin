package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var global_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List globally activated packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_listCmd).Standalone()

	global_listCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	globalCmd.AddCommand(global_listCmd)
}
