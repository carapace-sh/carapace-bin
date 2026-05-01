package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bCmd = &cobra.Command{
	Use:   "b",
	Short: "Benchmark",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bCmd).Standalone()

	addCommonFlags(bCmd)
	addCommonFlagCompletions(bCmd)
	rootCmd.AddCommand(bCmd)
}
