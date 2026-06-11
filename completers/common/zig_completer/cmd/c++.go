package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cxxCmd = &cobra.Command{
	Use:                "c++",
	Short:              "Use Zig as a drop-in C++ compiler",
	Run:                func(cmd *cobra.Command, args []string) {},
	DisableFlagParsing: true,
}

func init() {
	carapace.Gen(cxxCmd).Standalone()

	rootCmd.AddCommand(cxxCmd)

	carapace.Gen(cxxCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
