package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var analyzeProfileCmd = &cobra.Command{
	Use:   "analyze-profile",
	Short: "Analyzes build profile data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyzeProfileCmd).Standalone()

	analyzeProfileCmd.Flags().BoolP("dump", "d", false, "output full profile data dump either in human-readable 'text' format or script-friendly 'raw' format.")
	rootCmd.AddCommand(analyzeProfileCmd)
}
