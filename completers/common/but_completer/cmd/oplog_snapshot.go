package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var oplog_snapshotCmd = &cobra.Command{
	Use:   "snapshot",
	Short: "Create an on-demand snapshot with optional message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(oplog_snapshotCmd).Standalone()

	oplog_snapshotCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	oplog_snapshotCmd.Flags().StringP("message", "m", "", "Message to include with the snapshot")
	oplogCmd.AddCommand(oplog_snapshotCmd)
}
