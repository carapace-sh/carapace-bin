package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groupRemoveCmd = &cobra.Command{
	Use:   "remove [options] <group-spec>...",
	Short: "remove comps groups, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groupRemoveCmd).Standalone()

	groupRemoveCmd.Flags().Bool("no-packages", false, "Remove group without packages")
	groupRemoveCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	groupRemoveCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	groupRemoveCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	groupCmd.AddCommand(groupRemoveCmd)

	carapace.Gen(groupRemoveCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
