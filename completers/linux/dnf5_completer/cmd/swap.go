package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var swapCmd = &cobra.Command{
	Use:   "swap [options] <remove_spec> <install_spec>",
	Short: "remove software and install another in a single transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapCmd).Standalone()

	swapCmd.Flags().Bool("allow-erasing", false, "Allow erasing packages")
	swapCmd.Flags().String("from-repo", "", "Select items only from specified repositories")
	swapCmd.Flags().String("from-vendor", "", "Select items only from specified vendors")
	swapCmd.Flags().Bool("offline", false, "Store the transaction to be performed offline")
	swapCmd.Flags().String("store", "", "Store the current transaction in a directory at the specified path")
	swapCmd.Flags().Bool("transient", false, "Set up a transient overlay on /usr that will be discarded on reboot")

	rootCmd.AddCommand(swapCmd)

	carapace.Gen(swapCmd).FlagCompletion(carapace.ActionMap{
		"store": carapace.ActionDirectories(),
	})

	carapace.Gen(swapCmd).PositionalCompletion(
		action.ActionInstalledPackages(swapCmd),
		action.ActionPackageSearch(swapCmd),
	)
}
