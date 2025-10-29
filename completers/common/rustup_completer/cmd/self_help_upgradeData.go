package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var self_help_upgradeDataCmd = &cobra.Command{
	Use:   "upgrade-data",
	Short: "Upgrade the internal data format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(self_help_upgradeDataCmd).Standalone()

	self_helpCmd.AddCommand(self_help_upgradeDataCmd)
}
