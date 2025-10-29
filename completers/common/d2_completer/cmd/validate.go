package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates file.d2",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).PositionalCompletion(
		carapace.ActionFiles(".d2"),
	)
}
