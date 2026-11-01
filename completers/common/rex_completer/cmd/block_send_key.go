package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/rex"
	"github.com/spf13/cobra"
)

var blockSendKeyCmd = &cobra.Command{
	Use:   "send-key KEY...",
	Short: "Send key presses to a terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockSendKeyCmd).Standalone()

	blockSendKeyCmd.Flags().String("action", "", "send one ACTION for each key: press, release, or repeat")

	common.AddScopeFlags(blockSendKeyCmd)

	blockCmd.AddCommand(blockSendKeyCmd)

	carapace.Gen(blockSendKeyCmd).FlagCompletion(carapace.ActionMap{
		"action": carapace.ActionValues("press", "release", "repeat"),
	})

	carapace.Gen(blockSendKeyCmd).PositionalAnyCompletion(
		rex.ActionKeys(),
	)
}
