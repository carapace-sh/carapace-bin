package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_rmAllCmd = &cobra.Command{
	Use:   "rm-all",
	Short: "Remove multiple configuration values",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_rmAllCmd).Standalone()
	config_rmAllCmd.PersistentFlags().Bool("path", false, "Parse the keys as paths in a map or list rather than raw strings")
	configCmd.AddCommand(config_rmAllCmd)

	carapace.Gen(config_rmAllCmd).PositionalCompletion(
		action.ActionConfigKeys(config_rmAllCmd).FilterArgs(),
	)
}
