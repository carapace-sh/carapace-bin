package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search [pattern]...",
	Short: "search in package descriptions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("full", "f", false, "show full output")
	searchCmd.Flags().BoolP("names-only", "n", false, "show only names")
	rootCmd.AddCommand(searchCmd)
}
