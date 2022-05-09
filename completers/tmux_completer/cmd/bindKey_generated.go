package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var bindKeyCmd = &cobra.Command{
	Use:   "bind-key",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bindKeyCmd).Standalone()

	bindKeyCmd.Flags().StringS("N", "N", "", "note")
	bindKeyCmd.Flags().StringS("T", "T", "", "key-table")
	bindKeyCmd.Flags().BoolS("n", "n", false, "TODO description")
	bindKeyCmd.Flags().BoolS("r", "r", false, "TODO description")
	rootCmd.AddCommand(bindKeyCmd)
}
