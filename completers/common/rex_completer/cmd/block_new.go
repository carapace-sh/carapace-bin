package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/spf13/cobra"
)

var blockNewCmd = &cobra.Command{
	Use:   "new FLAVOR",
	Short: "Create a block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockNewCmd).Standalone()

	blockCmd.AddCommand(blockNewCmd)

	carapace.Gen(blockNewCmd).PositionalCompletion(
		action.ActionBlockFlavors(blockNewCmd),
	)
}
