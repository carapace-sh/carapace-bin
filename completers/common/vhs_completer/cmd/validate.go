package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate <file>...",
	Short: "Validate a glob file path and parses all the files to ensure they are valid without running them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	rootCmd.AddCommand(validateCmd)

	carapace.Gen(validateCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".tape"),
	)
}
