package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Access and manage snapshots",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_snapshotCmd).Standalone()
	computeCmd.AddCommand(compute_snapshotCmd)
}
