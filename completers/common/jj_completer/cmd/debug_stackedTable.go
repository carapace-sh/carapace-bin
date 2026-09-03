package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_stackedTableCmd = &cobra.Command{
	Use:   "stacked-table",
	Short: "Show stats of stacked table",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_stackedTableCmd).Standalone()

	debug_stackedTableCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debug_stackedTableCmd.Flags().StringP("key-size", "n", "", "Key size in bytes")
	debug_stackedTableCmd.MarkFlagRequired("key-size")
	debugCmd.AddCommand(debug_stackedTableCmd)

	carapace.Gen(debug_stackedTableCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}