package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var getwindowclassnameCmd = &cobra.Command{
	Use:   "getwindowclassname",
	Short: "Prints the class name for the window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowclassnameCmd).Standalone()

	rootCmd.AddCommand(getwindowclassnameCmd)

	carapace.Gen(getwindowclassnameCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
