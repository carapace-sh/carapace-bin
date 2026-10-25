package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/balooctl6_completer/cmd/action"
	"github.com/spf13/cobra"
)

var config_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a value to config parameter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_addCmd).Standalone()
	configCmd.AddCommand(config_addCmd)

	carapace.Gen(config_addCmd).PositionalCompletion(
		carapace.ActionValues("excludeFilters", "excludeFolders", "excludeMimetypes", "includeFolders"),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			switch strings.ToLower(c.Args[0]) {
			case "includefolders", "excludefolders":
				return carapace.ActionDirectories()
			case "excludefilters", "excludemimetypes":
				return action.ActionConfigValues(strings.ToLower(c.Args[0]))
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
