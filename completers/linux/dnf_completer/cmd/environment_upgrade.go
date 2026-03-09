package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environment_upgradeCmd = &cobra.Command{
	Use:   "upgrade [options] <group-spec|environment-spec>...",
	Short: "upgrade environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environment_upgradeCmd)

	environment_upgradeCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	environment_upgradeCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	environment_upgradeCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	environment_upgradeCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	environment_upgradeCmd.Flags().Bool("no-packages", false, "operate on environments without manipulating packages")
	environment_upgradeCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	environment_upgradeCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	environment_upgradeCmd.Flags().Bool("with-optional", false, "include optional packages")
	environmentCmd.AddCommand(environment_upgradeCmd)
}
