package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zig"
	"github.com/spf13/cobra"
)

var translateCCmd = &cobra.Command{
	Use:     "translate-c [options] [file]",
	Short:   "Convert C code to Zig code",
	GroupID: "source",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(translateCCmd).Standalone()

	common.AddBuildOptions(translateCCmd)
	rootCmd.AddCommand(translateCCmd)

	carapace.Gen(translateCCmd).PositionalCompletion(
		zig.ActionFiles(),
	)
}
