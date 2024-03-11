package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bunCmd = &cobra.Command{
	Use:   "bun",
	Short: "Bundle dependencies of input files into a .bun",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bunCmd).Standalone()

	rootCmd.AddCommand(bunCmd)

	carapace.Gen(bunCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
