package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var revUpgradeCmd = &cobra.Command{
	Use:   "rev-upgrade",
	Short: "Check for broken binaries and rebuild ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(revUpgradeCmd).Standalone()
	revUpgradeCmd.Flags().Bool("id-loadcmd-check", false, "Run id-loadcmd check")
	rootCmd.AddCommand(revUpgradeCmd)
}
