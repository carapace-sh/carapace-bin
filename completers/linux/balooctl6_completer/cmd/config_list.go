package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls", "show"},
	Short:   "Show the value of a config parameter",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()
	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).PositionalCompletion(
		carapace.ActionValues("contentIndexing", "excludeFilters", "excludeFolders", "excludeMimetypes", "hidden", "includeFolders"),
	)
}
