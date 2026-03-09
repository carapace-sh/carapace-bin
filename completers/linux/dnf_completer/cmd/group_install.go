package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupInstallCmd = &cobra.Command{
	Use:   "install [options] <group-spec>...",
	Short: "install groups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupInstallCmd).Standalone()

	groupInstallCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	groupInstallCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	groupInstallCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	groupInstallCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	groupInstallCmd.Flags().Bool("no-packages", false, "operate on groups without manipulating packages")
	groupInstallCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	groupInstallCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	groupInstallCmd.Flags().Bool("with-optional", false, "include optional packages")

	groupCmd.AddCommand(groupInstallCmd)
}
