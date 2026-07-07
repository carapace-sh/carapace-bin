package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete history entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().Bool("case-sensitive", false, "case-sensitive search")
	deleteCmd.Flags().BoolP("exact", "e", false, "exact match")
	deleteCmd.Flags().Bool("null", false, "NUL-terminated output")
	deleteCmd.Flags().BoolP("prefix", "p", false, "search by prefix")
	deleteCmd.Flags().BoolS("z", "z", false, "NUL-terminated output")

	rootCmd.AddCommand(deleteCmd)
}
