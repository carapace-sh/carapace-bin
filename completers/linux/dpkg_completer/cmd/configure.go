package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/apt"
	"github.com/spf13/cobra"
)

var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Reconfigure an unpacked package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configureCmd).Standalone()

	configureCmd.Flags().BoolP("pending", "a", false, "configure all unpacked but unconfigured packages")

	carapace.Gen(configureCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if configureCmd.Flag("pending").Changed {
				return carapace.ActionValues()
			}
			return apt.ActionPackages()
		}),
	)
}
