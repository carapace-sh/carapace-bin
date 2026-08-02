package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/dnf5_completer/cmd/action"
	"github.com/spf13/cobra"
)

var providesCmd = &cobra.Command{
	Use:   "provides [options] <package-spec>...",
	Short: "find what package provides the given value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(providesCmd).Standalone()

	rootCmd.AddCommand(providesCmd)

	carapace.Gen(providesCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(providesCmd),
	)
}
