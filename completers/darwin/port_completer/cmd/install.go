package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install and activate a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var reclaimCmd = &cobra.Command{
	Use:   "reclaim",
	Short: "Reclaim disk space by uninstalling inactive ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var revUpgradeCmd = &cobra.Command{
	Use:   "rev-upgrade",
	Short: "Check for broken binaries and rebuild ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var selectCmd = &cobra.Command{
	Use:   "select",
	Short: "Select a default version for a group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var setrequestedCmd = &cobra.Command{
	Use:   "setrequested",
	Short: "Mark a port as requested",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync the ports tree from the MacPorts rsync server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var selfupdateCmd = &cobra.Command{
	Use:   "selfupdate",
	Short: "Update MacPorts system, ports tree, and base tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var uninstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Deactivate and uninstall a port",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var unsetrequestedCmd = &cobra.Command{
	Use:   "unsetrequested",
	Short: "Mark a port as unrequested",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var upgradeCmd = &cobra.Command{
	Use:   "upgrade",
	Short: "Upgrade a port and its dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installCmd).Standalone()
	installCmd.Flags().Bool("no-rev-upgrade", false, "Do not run rev-upgrade after install")
	installCmd.Flags().Bool("unrequested", false, "Mark as unrequested (like a dependency)")
	rootCmd.AddCommand(installCmd)

	carapace.Gen(reclaimCmd).Standalone()
	reclaimCmd.Flags().Bool("disable-reminders", false, "Disable regular reclaim reminders")
	reclaimCmd.Flags().Bool("enable-reminders", false, "Enable regular reclaim reminders")
	rootCmd.AddCommand(reclaimCmd)

	carapace.Gen(revUpgradeCmd).Standalone()
	revUpgradeCmd.Flags().Bool("id-loadcmd-check", false, "Run id-loadcmd check")
	rootCmd.AddCommand(revUpgradeCmd)

	carapace.Gen(selectCmd).Standalone()
	selectCmd.Flags().Bool("list", false, "List available versions for a group")
	selectCmd.Flags().Bool("set", false, "Set a version for a group")
	selectCmd.Flags().Bool("show", false, "Show the current version for a group")
	selectCmd.Flags().Bool("summary", false, "Show a summary of all groups")
	rootCmd.AddCommand(selectCmd)

	carapace.Gen(setrequestedCmd).Standalone()
	rootCmd.AddCommand(setrequestedCmd)

	carapace.Gen(syncCmd).Standalone()
	rootCmd.AddCommand(syncCmd)

	carapace.Gen(selfupdateCmd).Standalone()
	selfupdateCmd.Flags().Bool("no-sync", false, "Only update MacPorts itself, do not update ports tree")
	rootCmd.AddCommand(selfupdateCmd)

	carapace.Gen(uninstallCmd).Standalone()
	uninstallCmd.Flags().Bool("follow-dependencies", false, "Also uninstall unneeded dependency ports")
	uninstallCmd.Flags().Bool("follow-dependents", false, "Recursively uninstall dependent ports first")
	uninstallCmd.Flags().Bool("no-exec", false, "Do not run uninstall hooks")
	rootCmd.AddCommand(uninstallCmd)

	carapace.Gen(unsetrequestedCmd).Standalone()
	rootCmd.AddCommand(unsetrequestedCmd)

	carapace.Gen(upgradeCmd).Standalone()
	upgradeCmd.Flags().Bool("enforce-variants", false, "Upgrade even if not outdated to change variants")
	upgradeCmd.Flags().Bool("force", false, "Always consider ports outdated")
	upgradeCmd.Flags().Bool("no-replace", false, "Do not install replacement ports")
	rootCmd.AddCommand(upgradeCmd)
}
