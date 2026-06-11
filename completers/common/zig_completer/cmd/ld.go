package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ldCmd = &cobra.Command{
	Use:                "ld",
	Short:              "Use Zig as a drop-in linker",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(ldCmd).Standalone()

	rootCmd.AddCommand(ldCmd)

	carapace.Gen(ldCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
