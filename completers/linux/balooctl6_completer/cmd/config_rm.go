package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_rmCmd = &cobra.Command{
	Use:     "rm",
	Short:   "Remove a value from a config option",
	Aliases: []string{"remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_rmCmd).Standalone()

	configCmd.AddCommand(config_rmCmd)

	carapace.Gen(config_rmCmd).PositionalCompletion(
		carapace.ActionValues(
			"includeFolders",
			"excludeFolders",
			"excludeFilters",
			"excludeMimetypes",
		),
	)
	carapace.Gen(config_rmCmd).PositionalAnyCompletion(
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
