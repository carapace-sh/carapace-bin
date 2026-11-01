package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/balooctl6_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a value from config parameter",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_removeCmd).Standalone()
	configCmd.AddCommand(config_removeCmd)

	carapace.Gen(config_removeCmd).PositionalCompletion(
		carapace.ActionValues("excludeFilters", "excludeFolders", "excludeMimetypes", "includeFolders"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionConfigValues(c.Args[0])
		}),
	)
}
