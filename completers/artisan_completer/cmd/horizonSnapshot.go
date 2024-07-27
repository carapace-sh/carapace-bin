package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonSnapshotCmd = &cobra.Command{
	Use:   "horizon:snapshot",
	Short: "Store a snapshot of the queue metrics",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonSnapshotCmd).Standalone()

	rootCmd.AddCommand(horizonSnapshotCmd)
}
