package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_removeCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a value from a configuration parameter",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_removeCmd).Standalone()
	configCmd.AddCommand(config_removeCmd)

	carapace.Gen(config_removeCmd).PositionalCompletion(
		carapace.ActionValues("includeFolders", "excludeFolders", "excludeFilters", "excludeMimetypes"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch c.Args[0] {
			case "includeFolders", "excludeFolders":
				return carapace.ActionDirectories()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
