package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "Show the value of a config option",
	Aliases: []string{"ls", "show"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_listCmd).Standalone()

	configCmd.AddCommand(config_listCmd)

	carapace.Gen(config_listCmd).PositionalAnyCompletion(
		carapace.ActionValues(
			"hidden",
			"contentIndexing",
			"includeFolders",
			"excludeFolders",
			"excludeFilters",
			"excludeMimetypes",
		),
	)
}
