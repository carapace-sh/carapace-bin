package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [options] [files] [-- args]",
	Short:   "Create executable and run immediately",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	common.AddBuildOptions(runCmd)
	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		zig.ActionFiles().FilterArgs(),
	)
}
