package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentCmd = &cobra.Command{
	Use:   "environment <subcommand> [options] [<environment-spec>...]",
	Short: "manage comps environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var environmentListCmd = &cobra.Command{
	Use:   "list [environment-spec]...",
	Short: "list environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var environmentInfoCmd = &cobra.Command{
	Use:   "info [environment-spec]...",
	Short: "print detailed information about environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var environmentInstallCmd = &cobra.Command{
	Use:   "install [options] <group-spec|environment-spec>...",
	Short: "install environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var environmentRemoveCmd = &cobra.Command{
	Use:   "remove [options] <group-spec|environment-spec>...",
	Short: "remove environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var environmentUpgradeCmd = &cobra.Command{
	Use:   "upgrade [options] <group-spec|environment-spec>...",
	Short: "upgrade environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentCmd).Standalone()

	environmentCmd.AddCommand(environmentListCmd)
	environmentCmd.AddCommand(environmentInfoCmd)
	environmentCmd.AddCommand(environmentInstallCmd)
	environmentCmd.AddCommand(environmentRemoveCmd)
	environmentCmd.AddCommand(environmentUpgradeCmd)

	for _, subcmd := range []*cobra.Command{environmentListCmd, environmentInfoCmd} {
		subcmd.Flags().Bool("available", false, "show only available environments")
		subcmd.Flags().Bool("installed", false, "show only installed environments")
	}

	for _, subcmd := range []*cobra.Command{environmentInstallCmd, environmentRemoveCmd, environmentUpgradeCmd} {
		subcmd.Flags().Bool("no-packages", false, "operate on environments without manipulating packages")
		subcmd.Flags().Bool("with-optional", false, "include optional packages")
	}

	for _, subcmd := range []*cobra.Command{environmentInstallCmd, environmentUpgradeCmd} {
		subcmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
		subcmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
		subcmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
		subcmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
		subcmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
		subcmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	}

	rootCmd.AddCommand(environmentCmd)
}
