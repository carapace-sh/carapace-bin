package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a value to a config option",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_addCmd).Standalone()

	configCmd.AddCommand(config_addCmd)

	carapace.Gen(config_addCmd).PositionalCompletion(
		carapace.ActionValues(
			"includeFolders",
			"excludeFolders",
			"excludeFilters",
			"excludeMimetypes",
		),
	)
	carapace.Gen(config_addCmd).PositionalAnyCompletion(
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
