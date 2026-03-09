package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupUpgradeCmd = &cobra.Command{
	Use:   "upgrade [options] <group-spec>...",
	Short: "upgrade groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupUpgradeCmd).Standalone()

	groupUpgradeCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	groupUpgradeCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	groupUpgradeCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	groupUpgradeCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	groupUpgradeCmd.Flags().Bool("no-packages", false, "operate on groups without manipulating packages")
	groupUpgradeCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	groupUpgradeCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	groupUpgradeCmd.Flags().Bool("with-optional", false, "include optional packages")

	groupCmd.AddCommand(groupUpgradeCmd)
}
