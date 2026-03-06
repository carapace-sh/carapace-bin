package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:   "autoremove",
	Short: "remove unneeded packages from the system",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	autoremoveCmd.Flags().Bool("offline", false, "store the transaction to be performed offline")
	autoremoveCmd.Flags().String("store", "", "store the current transaction in a directory")

	rootCmd.AddCommand(autoremoveCmd)
}
