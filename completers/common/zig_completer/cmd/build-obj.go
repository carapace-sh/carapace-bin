package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildObjCmd = &cobra.Command{
	Use:   "build-obj",
	Short: "Create object from source or object files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildObjCmd).Standalone()

	addBuildFlags(buildObjCmd)
	rootCmd.AddCommand(buildObjCmd)

	carapace.Gen(buildObjCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig", ".c", ".cpp", ".o", ".a"),
	)
}
