package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().String("manifest", "", "The path to the manifest to be validated")
	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).FlagCompletion(carapace.ActionMap{
		"manifest": carapace.ActionFiles(),
	})
}
