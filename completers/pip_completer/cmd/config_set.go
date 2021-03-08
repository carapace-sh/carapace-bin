package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pip_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_setCmd).Standalone()
	configCmd.AddCommand(config_setCmd)

	carapace.Gen(config_setCmd).PositionalCompletion(
		action.ActionConfigValues(config_setCmd),
	)
}
