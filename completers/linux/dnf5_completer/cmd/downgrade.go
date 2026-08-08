package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var downgradeCmd = &cobra.Command{
	Use:   "downgrade [options] <package-spec>...",
	Short: "downgrade software",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(downgradeCmd).Standalone()

	downgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies for resolve of requested operation")
	downgradeCmd.Flags().Bool("allowerasing", false, "Allow removing of installed packages to resolve problems")
	downgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	downgradeCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	downgradeCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	downgradeCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	downgradeCmd.Flags().Bool("no-allow-downgrade", false, "Disable downgrade of dependencies for resolve of requested operation")
	downgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	downgradeCmd.Flags().Bool("skip-broken", false, "Allow resolving of depsolve problems by skipping packages")
	downgradeCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")
	downgradeCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	downgradeCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(downgradeCmd)

	carapace.Gen(downgradeCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})

	carapace.Gen(downgradeCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(downgradeCmd),
	)
}
