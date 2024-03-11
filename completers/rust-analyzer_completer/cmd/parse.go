package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Parse stdin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(parseCmd).Standalone()

	parseCmd.Flags().Bool("no-dump", false, "Suppress printing")
	rootCmd.AddCommand(parseCmd)
}
