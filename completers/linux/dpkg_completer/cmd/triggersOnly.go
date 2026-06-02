package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var triggersOnlyCmd = &cobra.Command{
	Use:   "triggers-only",
	Short: "Processes only triggers",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(triggersOnlyCmd).Standalone()

	triggersOnlyCmd.Flags().BoolP("pending", "a", false, "configure all unpacked but unconfigured packages")

	carapace.Gen(triggersOnlyCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if triggersOnlyCmd.Flag("pending").Changed {
				return carapace.ActionValues()
			}
			return apt.ActionPackages()
		}),
	)
}
