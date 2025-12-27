package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apt_completer/cmd/common"
	"github.com/spf13/cobra"
)

var distUpgradeCmd = &cobra.Command{
	Use:   "dist-upgrade",
	Short: "upgrade the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distUpgradeCmd).Standalone()

	common.AddGetFlags(distUpgradeCmd)
	common.ActionInstallFlags(distUpgradeCmd)
	rootCmd.AddCommand(distUpgradeCmd)
}
