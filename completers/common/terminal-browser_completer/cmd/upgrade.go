package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade to the latest release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upgradeCmd).Standalone()

	rootCmd.AddCommand(upgradeCmd)
}
