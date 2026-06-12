package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/spf13/cobra"
)

var testObjCmd = &cobra.Command{
	Use:     "test-obj [options] [files]",
	Short:   "Create object for unit testing",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testObjCmd).Standalone()

	common.AddBuildOptions(testObjCmd)
	rootCmd.AddCommand(testObjCmd)

	common.AddTestOptions(testObjCmd)
	carapace.Gen(testObjCmd).PositionalAnyCompletion(
		carapace.ActionFiles().FilterArgs(),
	)
}
