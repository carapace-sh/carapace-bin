package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pruneHistoriesCmd = &cobra.Command{
	Use:   "prune-histories",
	Short: "clean up build histories",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pruneHistoriesCmd).Standalone()

	pruneHistoriesCmd.Flags().String("format", "", "Format the output using the given Go template")
	rootCmd.AddCommand(pruneHistoriesCmd)
}
