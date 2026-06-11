package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var arCmd = &cobra.Command{
	Use:                "ar",
	Short:              "Use Zig as a drop-in archiver",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(arCmd).Standalone()

	rootCmd.AddCommand(arCmd)

	carapace.Gen(arCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
