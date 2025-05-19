package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sanctumPruneExpiredCmd = &cobra.Command{
	Use:   "sanctum:prune-expired [--hours [HOURS]]",
	Short: "Prune tokens expired for more than specified number of hours",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sanctumPruneExpiredCmd).Standalone()

	sanctumPruneExpiredCmd.Flags().String("hours", "", "The number of hours to retain expired Sanctum tokens")
	rootCmd.AddCommand(sanctumPruneExpiredCmd)
}
