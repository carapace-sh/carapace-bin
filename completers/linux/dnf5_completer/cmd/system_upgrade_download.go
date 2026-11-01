package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgradeDownloadCmd = &cobra.Command{
	Use:   "download [options]",
	Short: "download everything needed to upgrade to a new release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgradeDownloadCmd).Standalone()

	systemUpgradeDownloadCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	systemUpgradeDownloadCmd.Flags().Bool("no-downgrade", false, "Do not install packages from the new release if they are older than what is currently installed")

	systemUpgradeCmd.AddCommand(systemUpgradeDownloadCmd)
}
