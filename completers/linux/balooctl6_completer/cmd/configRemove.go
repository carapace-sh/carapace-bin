package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"rm"},
	Short:   "Remove a value from a config parameter",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configRemoveCmd).Standalone()

	configCmd.AddCommand(configRemoveCmd)

	carapace.Gen(configRemoveCmd).PositionalCompletion(
		actionConfigModifyParameters(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			switch c.Args[0] {
			case "includeFolders", "excludeFolders":
				return carapace.ActionDirectories()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
