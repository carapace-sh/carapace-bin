package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var blockCallCmd = &cobra.Command{
	Use:   "call METHOD [ARGS]",
	Short: "Call a block method",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockCallCmd).Standalone()

	common.AddScopeFlags(blockCallCmd)

	blockCmd.AddCommand(blockCallCmd)

	carapace.Gen(blockCallCmd).PositionalCompletion(
		rex.ActionApiMethods(),
		carapace.Batch(
			carapace.ActionFiles().Prefix("@"),
			carapace.ActionValues("-"),
		).ToA(),
	)
}
