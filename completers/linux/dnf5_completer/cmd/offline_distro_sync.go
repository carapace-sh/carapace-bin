package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineDistroSyncCmd = &cobra.Command{
	Use:   "offline-distro-sync [subcommand]",
	Short: "store a distro-sync transaction to be performed offline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineDistroSyncCmd).Standalone()

	rootCmd.AddCommand(offlineDistroSyncCmd)
}
