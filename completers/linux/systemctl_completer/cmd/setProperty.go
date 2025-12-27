package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/systemctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setPropertyCmd = &cobra.Command{
	Use:     "set-property",
	Short:   "Sets one or more properties of a unit",
	GroupID: "unit",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setPropertyCmd).Standalone()

	rootCmd.AddCommand(setPropertyCmd)

	carapace.Gen(setPropertyCmd).PositionalCompletion(
		action.ActionUnits(setPropertyCmd),
	)

	carapace.Gen(setPropertyCmd).PositionalAnyCompletion(
		carapace.ActionMultiParts("=", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				return action.ActionProperties(setPropertyCmd, c.Args[0]).Invoke(c).Suffix("=").ToA()
			default:
				return carapace.ActionValues()
			}
		}),
	)
}
