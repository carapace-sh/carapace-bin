package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "remove a configuration setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_rmCmd).Standalone()
	configCmd.AddCommand(config_rmCmd)

	carapace.Gen(config_rmCmd).PositionalCompletion(
		action.ActionConfigKeys(),
	)
}
