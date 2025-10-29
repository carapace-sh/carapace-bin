package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_self_upgradeDataCmd = &cobra.Command{
	Use:   "upgrade-data",
	Short: "Upgrade the internal data format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_self_upgradeDataCmd).Standalone()

	help_selfCmd.AddCommand(help_self_upgradeDataCmd)
}
