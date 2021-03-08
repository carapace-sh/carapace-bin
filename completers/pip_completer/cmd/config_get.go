package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pip_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the value associated with name",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_getCmd).Standalone()
	configCmd.AddCommand(config_getCmd)

	carapace.Gen(config_getCmd).PositionalCompletion(
		action.ActionConfigValues(config_getCmd),
	)
}
