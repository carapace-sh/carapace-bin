package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var backfillCmd = &cobra.Command{
	Use:     "backfill",
	Short:   "Fetch missing objects in a partial clone on demand",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(backfillCmd).Standalone()

	backfillCmd.Flags().String("min-batch-size", "", "Minimum number of objects to request at a time")
	backfillCmd.Flags().Bool("sparse", false, "Restrict the missing objects to the current sparse-checkout")
	rootCmd.AddCommand(backfillCmd)
}
