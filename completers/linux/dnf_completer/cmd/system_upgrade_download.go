package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_upgradeDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download packages needed for upgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_upgradeDownloadCmd).Standalone()

	system_upgradeDownloadCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	system_upgradeDownloadCmd.Flags().Bool("no-downgrade", false, "do not install older packages")
	system_upgradeDownloadCmd.Flags().String("releasever", "", "required: the version to upgrade to")

	system_upgradeCmd.AddCommand(system_upgradeDownloadCmd)
}
