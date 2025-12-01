package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snapshotCmd = &cobra.Command{
	Use:     "snapshot",
	Short:   "Create an on-demand snapshot with optional message",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "operation history",
}

func init() {
	carapace.Gen(snapshotCmd).Standalone()

	snapshotCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	snapshotCmd.Flags().StringP("message", "m", "", "Message to include with the snapshot")
	rootCmd.AddCommand(snapshotCmd)
}
