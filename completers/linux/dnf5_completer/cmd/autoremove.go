package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:   "autoremove",
	Short: "remove all unneeded packages originally installed as dependencies",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	autoremoveCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	autoremoveCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	autoremoveCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(autoremoveCmd)

	carapace.Gen(autoremoveCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
