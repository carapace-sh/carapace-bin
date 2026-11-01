package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistroSyncLogCmd = &cobra.Command{
	Use:   "log [options]",
	Short: "show logs from past offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistroSyncLogCmd).Standalone()

	offlineDistroSyncLogCmd.Flags().String("number", "", "Which log to show")

	offlineDistroSyncCmd.AddCommand(offlineDistroSyncLogCmd)
}
