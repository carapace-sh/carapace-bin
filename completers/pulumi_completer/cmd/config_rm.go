package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/pulumi_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_rmCmd = &cobra.Command{
	Use:   "rm",
	Short: "Remove configuration value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	config_rmCmd.PersistentFlags().Bool("path", false, "The key contains a path to a property in a map or list to remove")
	configCmd.AddCommand(config_rmCmd)

	carapace.Gen(config_rmCmd).PositionalCompletion(
		action.ActionConfigKeys(config_rmCmd),
	)
}
