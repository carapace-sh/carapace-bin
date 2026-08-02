package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var environmentRemoveCmd = &cobra.Command{
	Use:   "remove [options] <environment-spec>...",
	Short: "remove comps environments, including their packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(environmentRemoveCmd).Standalone()

	environmentRemoveCmd.Flags().Bool("no-packages", false, "Remove environment without packages")
	environmentRemoveCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	environmentRemoveCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	environmentRemoveCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	environmentCmd.AddCommand(environmentRemoveCmd)

	carapace.Gen(environmentRemoveCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
