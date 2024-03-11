package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove an installed Package including conffiles",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	purgeCmd.Flags().BoolP("pending", "a", false, "configure all unpacked but unconfigured packages")

	carapace.Gen(purgeCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if purgeCmd.Flag("pending").Changed {
				return carapace.ActionValues()
			}
			return apt.ActionPackages()
		}),
	)
}
