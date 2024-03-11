package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var store_signCmd = &cobra.Command{
	Use:   "sign",
	Short: "sign store paths",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_signCmd).Standalone()

	store_signCmd.Flags().BoolP("key-file", "k", false, "File containing the secret signing key")
	storeCmd.AddCommand(store_signCmd)

	addEvaluationFlags(store_signCmd)
	addFlakeFlags(store_signCmd)
	addLoggingFlags(store_signCmd)

	carapace.Gen(store_signCmd).FlagCompletion(carapace.ActionMap{
		"key-file": carapace.ActionFiles(),
	})

	// TODO positional completion
}
