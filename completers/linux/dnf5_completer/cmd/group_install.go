package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupInstallCmd = &cobra.Command{
	Use:   "install [options] <group-spec>...",
	Short: "install comps groups, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupInstallCmd).Standalone()

	groupInstallCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies")
	groupInstallCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	groupInstallCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	groupInstallCmd.Flags().Bool("no-packages", false, "Install group without packages")
	groupInstallCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	groupInstallCmd.Flags().Bool("skip-broken", false, "Skip broken packages")
	groupInstallCmd.Flags().Bool("skip-unavailable", false, "Skip unavailable packages")
	groupInstallCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	groupInstallCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")
	groupInstallCmd.Flags().Bool("with-optional", false, "Include optional packages")

	groupCmd.AddCommand(groupInstallCmd)

	carapace.Gen(groupInstallCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
