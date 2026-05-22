package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildLibCmd = &cobra.Command{
	Use:   "build-lib",
	Short: "Create library from source or object files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildLibCmd).Standalone()

	addBuildFlags(buildLibCmd)
	rootCmd.AddCommand(buildLibCmd)

	carapace.Gen(buildLibCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig", ".c", ".cpp", ".o", ".a"),
	)
}
