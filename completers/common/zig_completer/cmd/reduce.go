package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/zig_completer/cmd/common"
	"github.com/spf13/cobra"
)

var reduceCmd = &cobra.Command{
	Use:     "reduce [options] [file]",
	Short:   "Minimize a bug report",
	GroupID: "source",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reduceCmd).Standalone()

	common.AddBuildOptions(reduceCmd)
	rootCmd.AddCommand(reduceCmd)

	carapace.Gen(reduceCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".zig"),
	)
}
