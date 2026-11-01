package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var offlineUpgradeCmd = &cobra.Command{
	Use:   "offline-upgrade [subcommand]",
	Short: "store an upgrade transaction to be performed offline",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(offlineUpgradeCmd).Standalone()

	rootCmd.AddCommand(offlineUpgradeCmd)
}
