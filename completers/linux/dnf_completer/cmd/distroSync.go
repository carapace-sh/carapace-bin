package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var distroSyncCmd = &cobra.Command{
	Use:     "distro-sync [options] [<package-spec>...]",
	Aliases: []string{"distro-sync"},
	Short:   "upgrade or downgrade installed packages to the latest available version",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distroSyncCmd).Standalone()

	distroSyncCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")
	distroSyncCmd.Flags().Bool("downloadonly", false, "download packages without executing RPM transaction")
	distroSyncCmd.Flags().Bool("skip-broken", false, "resolve dependency problems by removing problematic packages")
	distroSyncCmd.Flags().Bool("skip-unavailable", false, "allow skipping packages that are not possible to synchronize")

	rootCmd.AddCommand(distroSyncCmd)

	carapace.Gen(distroSyncCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(distroSyncCmd),
	)
}
