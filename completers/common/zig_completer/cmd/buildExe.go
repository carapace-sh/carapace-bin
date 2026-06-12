package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/spf13/cobra"
)

var buildExeCmd = &cobra.Command{
	Use:     "build-exe [options] [files]",
	Short:   "Create executable from source or object files",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildExeCmd).Standalone()

	common.AddBuildOptions(buildExeCmd)
	rootCmd.AddCommand(buildExeCmd)

	carapace.Gen(buildExeCmd).PositionalAnyCompletion(
		zig.ActionFiles().FilterArgs(),
	)
}
