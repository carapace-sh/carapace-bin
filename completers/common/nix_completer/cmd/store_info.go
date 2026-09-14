package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/nix_completer/cmd/common"
	"github.com/spf13/cobra"
)

var store_infoCmd = &cobra.Command{
	Use:     "info",
	Short:   "test whether a store can be accessed",
	Aliases: []string{"ping"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_infoCmd).Standalone()

	common.AddJSONFlags(store_infoCmd)
	common.AddLoggingFlags(store_infoCmd)

	storeCmd.AddCommand(store_infoCmd)
}
