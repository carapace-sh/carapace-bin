package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate terraform code for stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()

	generateCmd.Flags().Bool("detailed-exit-code", false, "Return detailed exit code (0 = ok, 1 = errors, 2 = no errors but changes were made")
	rootCmd.AddCommand(generateCmd)
}
