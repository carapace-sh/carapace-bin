package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/spf13/cobra"
)

var buildObjCmd = &cobra.Command{
	Use:     "build-obj [options] [files]",
	Short:   "Create object from source or object files",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildObjCmd).Standalone()

	common.AddBuildOptions(buildObjCmd)
	rootCmd.AddCommand(buildObjCmd)

	carapace.Gen(buildObjCmd).PositionalAnyCompletion(
		zig.ActionFiles().FilterArgs(),
	)
}
