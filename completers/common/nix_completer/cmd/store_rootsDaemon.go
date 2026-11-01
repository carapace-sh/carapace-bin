package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var store_rootsDaemonCmd = &cobra.Command{
	Use:   "roots-daemon",
	Short: "run a daemon that returns garbage collector roots on request",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_rootsDaemonCmd).Standalone()

	storeCmd.AddCommand(store_rootsDaemonCmd)

	common.AddLoggingFlags(store_rootsDaemonCmd)
}
