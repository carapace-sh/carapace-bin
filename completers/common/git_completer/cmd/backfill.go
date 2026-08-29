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

	backfillCmd.Flags().Bool("include-edges", false, "request objects at the edges of the partial clone")
	backfillCmd.Flags().String("min-batch-size", "", "minimum number of objects to request at a time")
	backfillCmd.Flags().Bool("no-include-edges", false, "do not request objects at the edges of the partial clone")
	backfillCmd.Flags().Bool("no-sparse", false, "do not restrict the missing objects to the current sparse-checkout")
	backfillCmd.Flags().Bool("sparse", false, "restrict the missing objects to the current sparse-checkout")
	rootCmd.AddCommand(backfillCmd)
}
