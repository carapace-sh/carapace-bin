package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowlowerCmd = &cobra.Command{
	Use:   "windowlower",
	Short: "Lower a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowlowerCmd).Standalone()

	rootCmd.AddCommand(windowlowerCmd)

	carapace.Gen(windowlowerCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
