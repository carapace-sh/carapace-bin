package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var collectCmd = &cobra.Command{
	Use:   "collect",
	Short: "Collect input into a single output argument",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(collectCmd).Standalone()

	collectCmd.Flags().BoolS("N", "N", false, "don't trim trailing newlines")
	collectCmd.Flags().BoolP("allow-empty", "a", false, "always prints one argument")
	collectCmd.Flags().Bool("no-trim-newlines", false, "don't trim trailing newlines")

	rootCmd.AddCommand(collectCmd)
}
