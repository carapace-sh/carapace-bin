package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test [options] [files]",
	Short:   "Perform unit testing",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	common.AddBuildOptions(testCmd)
	rootCmd.AddCommand(testCmd)

	common.AddTestOptions(testCmd)
	carapace.Gen(testCmd).PositionalAnyCompletion(
		zig.ActionFiles().FilterArgs(),
	)
}
