package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var store_copySigsCmd = &cobra.Command{
	Use:   "copy-sigs",
	Short: "copy store path signatures from substituters",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_copySigsCmd).Standalone()

	store_copySigsCmd.Flags().StringP("substituter", "s", "", "Copy signatures from the specified store")
	storeCmd.AddCommand(store_copySigsCmd)

	addEvaluationFlags(store_copySigsCmd)
	addFlakeFlags(store_copySigsCmd)
	addLoggingFlags(store_copySigsCmd)

	// TODO positional completion

	// TODO positional completion
}
