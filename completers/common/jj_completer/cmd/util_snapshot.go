package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Snapshot the working copy if needed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_snapshotCmd).Standalone()

	util_snapshotCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_snapshotCmd)
}
