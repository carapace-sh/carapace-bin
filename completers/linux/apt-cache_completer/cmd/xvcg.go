package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apt-cache_completer/cmd/action"
	"github.com/spf13/cobra"
)

var xvcgCmd = &cobra.Command{
	Use:   "xvcg",
	Short: "generates output suitable for use by xvcg",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(xvcgCmd).Standalone()

	rootCmd.AddCommand(xvcgCmd)

	carapace.Gen(xvcgCmd).PositionalAnyCompletion(
		action.ActionPackages(),
	)
}
