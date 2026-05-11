package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a value to config parameter",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configAddCmd).Standalone()

	configCmd.AddCommand(configAddCmd)

	carapace.Gen(configAddCmd).PositionalCompletion(
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
