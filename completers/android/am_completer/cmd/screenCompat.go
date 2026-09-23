package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var screenCompatCmd = &cobra.Command{
	Use:   "screen-compat",
	Short: "Control screen compatibility mode of a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(screenCompatCmd).Standalone()

	rootCmd.AddCommand(screenCompatCmd)

	carapace.Gen(screenCompatCmd).PositionalCompletion(
		carapace.ActionValues("on", "off"),
		android.ActionPackages(),
	)
}
