package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/spf13/cobra"
)

var buildLibCmd = &cobra.Command{
	Use:     "build-lib [options] [files]",
	Short:   "Create library from source or object files",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildLibCmd).Standalone()

	common.AddBuildOptions(buildLibCmd)
	rootCmd.AddCommand(buildLibCmd)

	carapace.Gen(buildLibCmd).PositionalAnyCompletion(
		zig.ActionFiles().FilterArgs(),
	)
}
