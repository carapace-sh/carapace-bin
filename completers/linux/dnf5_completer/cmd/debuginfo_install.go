package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var debuginfoInstallCmd = &cobra.Command{
	Use:   "debuginfo-install [options] <package-spec>...",
	Short: "install debuginfo packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuginfoInstallCmd).Standalone()

	debuginfoInstallCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	debuginfoInstallCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	debuginfoInstallCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	debuginfoInstallCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	debuginfoInstallCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	debuginfoInstallCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(debuginfoInstallCmd)

	carapace.Gen(debuginfoInstallCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})

	carapace.Gen(debuginfoInstallCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(debuginfoInstallCmd),
	)
}
