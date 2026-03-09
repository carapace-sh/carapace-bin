package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_upgradeCmd = &cobra.Command{
	Use:   "system-upgrade <subcommand> [options]",
	Short: "upgrade the system to a new major release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_upgradeCmd).Standalone()

	rootCmd.AddCommand(system_upgradeCmd)
}
