package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debuginfoInstallCmd = &cobra.Command{
	Use:   "debuginfo-install [options] <package-spec>...",
	Short: "install the associated debuginfo packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuginfoInstallCmd).Standalone()

	debuginfoInstallCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	debuginfoInstallCmd.Flags().Bool("skip-broken", false, "resolve dependency problems by removing problematic packages")
	debuginfoInstallCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not available")

	rootCmd.AddCommand(debuginfoInstallCmd)

	carapace.Gen(debuginfoInstallCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(debuginfoInstallCmd),
	)
}
