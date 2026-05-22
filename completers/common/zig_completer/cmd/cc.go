package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ccCmd = &cobra.Command{
	Use:                "cc",
	Short:              "Use Zig as a drop-in C compiler",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(ccCmd).Standalone()

	rootCmd.AddCommand(ccCmd)

	carapace.Gen(ccCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
