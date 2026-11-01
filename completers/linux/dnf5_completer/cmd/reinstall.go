package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var reinstallCmd = &cobra.Command{
	Use:   "reinstall [options] <package-spec>...",
	Short: "reinstall software",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reinstallCmd).Standalone()

	reinstallCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	reinstallCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	reinstallCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	reinstallCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	reinstallCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	reinstallCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	reinstallCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	reinstallCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	reinstallCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	reinstallCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	reinstallCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	reinstallCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(reinstallCmd)

	carapace.Gen(reinstallCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})

	carapace.Gen(reinstallCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(reinstallCmd),
	)
}
