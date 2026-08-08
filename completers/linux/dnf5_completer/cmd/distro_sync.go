package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var distroSyncCmd = &cobra.Command{
	Use:   "distro-sync [options] [<package-spec>...]",
	Short: "upgrade or downgrade installed software to the latest available versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(distroSyncCmd).Standalone()

	distroSyncCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	distroSyncCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	distroSyncCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	distroSyncCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	distroSyncCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	distroSyncCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	distroSyncCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	distroSyncCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	distroSyncCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	distroSyncCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(distroSyncCmd)

	carapace.Gen(distroSyncCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})

	carapace.Gen(distroSyncCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(distroSyncCmd),
	)
}
