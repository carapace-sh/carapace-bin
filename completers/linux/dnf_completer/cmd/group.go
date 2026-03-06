package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupCmd = &cobra.Command{
	Use:   "group <subcommand> [options] [<group-spec>...]",
	Short: "manage comps groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var groupListCmd = &cobra.Command{
	Use:   "list [group-spec]...",
	Short: "list groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var groupInfoCmd = &cobra.Command{
	Use:   "info [group-spec]...",
	Short: "print detailed information about groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var groupInstallCmd = &cobra.Command{
	Use:   "install [options] <group-spec>...",
	Short: "install groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var groupRemoveCmd = &cobra.Command{
	Use:   "remove [options] <group-spec>...",
	Short: "remove groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var groupUpgradeCmd = &cobra.Command{
	Use:   "upgrade [options] <group-spec>...",
	Short: "upgrade groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupCmd).Standalone()

	groupCmd.AddCommand(groupListCmd)
	groupCmd.AddCommand(groupInfoCmd)
	groupCmd.AddCommand(groupInstallCmd)
	groupCmd.AddCommand(groupRemoveCmd)
	groupCmd.AddCommand(groupUpgradeCmd)

	for _, subcmd := range []*cobra.Command{groupListCmd, groupInfoCmd} {
		subcmd.Flags().Bool("available", false, "show only available groups")
		subcmd.Flags().Bool("installed", false, "show only installed groups")
		subcmd.Flags().Bool("hidden", false, "show also hidden groups")
		subcmd.Flags().String("contains-pkgs", "", "show only groups containing packages")
	}

	for _, subcmd := range []*cobra.Command{groupInstallCmd, groupRemoveCmd, groupUpgradeCmd} {
		subcmd.Flags().Bool("with-optional", false, "include optional packages")
		subcmd.Flags().Bool("no-packages", false, "operate on groups without manipulating packages")
	}

	for _, subcmd := range []*cobra.Command{groupInstallCmd, groupUpgradeCmd} {
		subcmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
		subcmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
		subcmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
		subcmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
		subcmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
		subcmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	}

	rootCmd.AddCommand(groupCmd)
}
