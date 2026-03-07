package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var systemUpgradeCmd = &cobra.Command{
	Use:   "system-upgrade <subcommand> [options]",
	Short: "upgrade the system to a new major release",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var systemUpgradeCleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove stored offline transaction and delete cached package files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var systemUpgradeDownloadCmd = &cobra.Command{
	Use:   "download",
	Short: "download packages needed for upgrade",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var systemUpgradeLogCmd = &cobra.Command{
	Use:   "log",
	Short: "see logs from offline transactions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var systemUpgradeRebootCmd = &cobra.Command{
	Use:   "reboot",
	Short: "prepare system to perform offline transaction and reboot",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(systemUpgradeCmd).Standalone()

	systemUpgradeCmd.AddCommand(systemUpgradeCleanCmd)
	systemUpgradeCmd.AddCommand(systemUpgradeDownloadCmd)
	systemUpgradeCmd.AddCommand(systemUpgradeLogCmd)
	systemUpgradeCmd.AddCommand(systemUpgradeRebootCmd)

	systemUpgradeDownloadCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	systemUpgradeDownloadCmd.Flags().Bool("no-downgrade", false, "do not install older packages")
	systemUpgradeDownloadCmd.Flags().String("releasever", "", "required: the version to upgrade to")

	systemUpgradeRebootCmd.Flags().Int("number", 0, "boot number")
	systemUpgradeLogCmd.Flags().Int("number", 0, "boot number")
	systemUpgradeRebootCmd.Flags().Bool("poweroff", false, "power off after transaction instead of rebooting")

	rootCmd.AddCommand(systemUpgradeCmd)
}
