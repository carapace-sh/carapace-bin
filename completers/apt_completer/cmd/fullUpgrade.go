package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apt_completer/cmd/common"
	"github.com/spf13/cobra"
)

var fullUpgradeCmd = &cobra.Command{
	Use:   "full-upgrade",
	Short: "upgrade the system including removal of packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fullUpgradeCmd).Standalone()

	common.AddGetFlags(fullUpgradeCmd)
	common.ActionInstallFlags(fullUpgradeCmd)
	rootCmd.AddCommand(fullUpgradeCmd)
}
