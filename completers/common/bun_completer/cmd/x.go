package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/npm"
	"github.com/spf13/cobra"
)

var xCmd = &cobra.Command{
	Use:   "x",
	Short: "Install and execute a package bin",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xCmd).Standalone()

	rootCmd.AddCommand(xCmd)

	carapace.Gen(xCmd).PositionalCompletion(
		npm.ActionPackageSearch(""),
	)
}
