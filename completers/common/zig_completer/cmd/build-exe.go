package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildExeCmd = &cobra.Command{
	Use:   "build-exe",
	Short: "Create executable from source or object files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildExeCmd).Standalone()

	addBuildFlags(buildExeCmd)
	rootCmd.AddCommand(buildExeCmd)

	carapace.Gen(buildExeCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig", ".c", ".cpp", ".o", ".a"),
	)
}
