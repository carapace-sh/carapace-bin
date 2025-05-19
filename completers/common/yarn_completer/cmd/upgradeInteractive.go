package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeInteractiveCmd = &cobra.Command{
	Use:     "upgrade-interactive",
	Short:   "Open the upgrade interface",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeInteractiveCmd).Standalone()

	rootCmd.AddCommand(upgradeInteractiveCmd)
}
