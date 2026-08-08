package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentUpgradeCmd = &cobra.Command{
	Use:   "upgrade [options] <environment-spec>...",
	Short: "upgrade comps environments, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentUpgradeCmd).Standalone()

	environmentUpgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies")
	environmentUpgradeCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	environmentUpgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	environmentUpgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	environmentUpgradeCmd.Flags().Bool("skip-unavailable", false, "Skip unavailable packages")
	environmentUpgradeCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	environmentUpgradeCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	environmentCmd.AddCommand(environmentUpgradeCmd)

	carapace.Gen(environmentUpgradeCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
