package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var blockSendCmd = &cobra.Command{
	Use:   "send [TEXT]",
	Short: "Send input to a terminal",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(blockSendCmd).Standalone()

	common.AddScopeFlags(blockSendCmd)

	blockCmd.AddCommand(blockSendCmd)
}
