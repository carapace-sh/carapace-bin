package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Echo the config values to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		action.ActionConfigKeys(getCmd),
	)
}
