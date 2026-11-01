package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var resizeCmd = &cobra.Command{
	Use:   "resize [BLOCK]",
	Short: "Resize a block",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resizeCmd).Standalone()

	common.AddWindowScopeFlags(resizeCmd)

	rootCmd.AddCommand(resizeCmd)

	carapace.Gen(resizeCmd).PositionalCompletion(
		action.ActionBlocks(resizeCmd),
	)
}
