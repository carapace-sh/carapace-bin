package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debuginfo_installCmd = &cobra.Command{
	Use:   "debuginfo-install [options] <package-spec>...",
	Short: "install the associated debuginfo packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuginfo_installCmd).Standalone()

	debuginfo_installCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	debuginfo_installCmd.Flags().Bool("skip-broken", false, "resolve dependency problems by removing problematic packages")
	debuginfo_installCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not available")

	rootCmd.AddCommand(debuginfo_installCmd)

	carapace.Gen(debuginfo_installCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(debuginfo_installCmd),
	)
}
