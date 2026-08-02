package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var markCmd = &cobra.Command{
	Use:   "mark [subcommand]",
	Short: "change the reason of an installed package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(markCmd).Standalone()

	markCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	markCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	markCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(markCmd)

	carapace.Gen(markCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})
}
