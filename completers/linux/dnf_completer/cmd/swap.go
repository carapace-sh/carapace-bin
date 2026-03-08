package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var swapCmd = &cobra.Command{
	Use:   "swap <remove-spec> <install-spec>",
	Short: "remove software and install another in a single transaction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swapCmd).Standalone()

	swapCmd.Flags().Bool("allowerasing", false, "allow removing of installed packages")

	rootCmd.AddCommand(swapCmd)

	carapace.Gen(swapCmd).PositionalCompletion(
		action.ActionInstalledPackages(swapCmd),
		action.ActionPackageSearch(swapCmd),
	)
}
