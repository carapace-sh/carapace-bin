package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var svidCmd = &cobra.Command{
	Use:    "svid",
	Short:  "Manage Teleport Workload Identity SVIDs.",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(svidCmd).Standalone()

	rootCmd.AddCommand(svidCmd)
}
