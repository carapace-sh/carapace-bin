package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/npm_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Sets each of the config keys to the value provided",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()
	setCmd.Flags().BoolP("global", "g", false, "operate in global mode")
	setCmd.Flags().String("location", "", "location of the config file to set in")

	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).FlagCompletion(carapace.ActionMap{
		"location": carapace.ActionValues("user", "global", "project"),
	})

	carapace.Gen(setCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionConfigKeys(setCmd).Invoke(c).Suffix("=").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
