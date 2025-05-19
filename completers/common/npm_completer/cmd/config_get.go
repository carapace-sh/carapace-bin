package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Echo the config values to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()

	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalAnyCompletion(
		action.ActionConfigKeys(config_getCmd),
	)
}
