package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupUpgradeCmd = &cobra.Command{
	Use:   "upgrade [options] <group-spec>...",
	Short: "upgrade comps groups, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupUpgradeCmd).Standalone()

	groupUpgradeCmd.Flags().Bool("allow-downgrade", false, "Allow downgrade of dependencies")
	groupUpgradeCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	groupUpgradeCmd.Flags().Bool("downloadonly", false, "Only download packages for a transaction")
	groupUpgradeCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	groupUpgradeCmd.Flags().Bool("skip-unavailable", false, "Skip unavailable packages")
	groupUpgradeCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	groupUpgradeCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	groupCmd.AddCommand(groupUpgradeCmd)

	carapace.Gen(groupUpgradeCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
