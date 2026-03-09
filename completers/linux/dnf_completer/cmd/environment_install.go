package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environment_installCmd = &cobra.Command{
	Use:   "install [options] <group-spec|environment-spec>...",
	Short: "install environments",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environment_installCmd)

	environment_installCmd.Flags().Bool("allow-downgrade", false, "enable downgrade of dependencies")
	environment_installCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	environment_installCmd.Flags().Bool("downloadonly", false, "download packages without executing transaction")
	environment_installCmd.Flags().Bool("no-allow-downgrade", false, "disable downgrade of dependencies")
	environment_installCmd.Flags().Bool("no-packages", false, "operate on environments without manipulating packages")
	environment_installCmd.Flags().Bool("skip-broken", false, "resolve dependency problems")
	environment_installCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages")
	environment_installCmd.Flags().Bool("with-optional", false, "include optional packages")
	environmentCmd.AddCommand(environment_installCmd)
}
