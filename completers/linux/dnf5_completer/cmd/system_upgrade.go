package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgradeCmd = &cobra.Command{
	Use:   "system-upgrade [subcommand]",
	Short: "prepare system for upgrade to a new release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgradeCmd).Standalone()

	rootCmd.AddCommand(systemUpgradeCmd)
}
