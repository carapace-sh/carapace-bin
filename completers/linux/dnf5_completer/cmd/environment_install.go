package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentInstallCmd = &cobra.Command{
	Use:   "install [options] <environment-spec>...",
	Short: "install comps environments, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentInstallCmd).Standalone()

	environmentInstallCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies")
	environmentInstallCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	environmentInstallCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	environmentInstallCmd.Flags().Bool("no-packages", false, "Install environment without packages")
	environmentInstallCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	environmentInstallCmd.Flags().Bool("skip-broken", false, "Skip broken packages")
	environmentInstallCmd.Flags().Bool("skip-unavailable", false, "Skip unavailable packages")
	environmentInstallCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	environmentInstallCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")
	environmentInstallCmd.Flags().Bool("with-optional", false, "Include optional packages")

	environmentCmd.AddCommand(environmentInstallCmd)

	carapace.Gen(environmentInstallCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
