package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var ddtpruneCmd = &cobra.Command{
	Use:     "ddtprune -d days|-p percentage pool",
	Short:   "prune dedup table entries",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ddtpruneCmd).Standalone()

	ddtpruneCmd.Flags().StringS("d", "d", "", "prune entries older than days")
	ddtpruneCmd.Flags().StringS("p", "p", "", "prune percentage of entries")

	rootCmd.AddCommand(ddtpruneCmd)

	carapace.Gen(ddtpruneCmd).PositionalCompletion(
		zfs.ActionPools(),
	)
}
