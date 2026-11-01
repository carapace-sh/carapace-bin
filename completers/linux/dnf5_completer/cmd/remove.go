package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:     "remove [options] <package-spec>|@<group-spec>|@<environment-spec>...",
	Aliases: []string{"rm"},
	Short:   "remove (uninstall) software",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removeCmd).Standalone()

	removeCmd.Flags().Bool("duplicates", false, "Remove older versions of duplicate packages")
	removeCmd.Flags().String("installed-from-repo", "", "Filter installed packages by repository ID")
	removeCmd.Flags().Int("limit", 0, "Limit the number of installonly package versions to keep")
	removeCmd.Flags().Bool("no-autoremove", false, "Disable removal of dependencies that are no longer used")
	removeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	removeCmd.Flags().Bool("oldinstallonly", false, "Remove old installonly packages")
	removeCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	removeCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(removeCmd)

	carapace.Gen(removeCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})

	carapace.Gen(removeCmd).PositionalAnyCompletion(
		action.ActionInstalledPackages(removeCmd),
	)
}
