package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var getDesktopForWindowCmd = &cobra.Command{
	Use:   "get_desktop_for_window",
	Short: "Output the desktop currently containing the given window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getDesktopForWindowCmd).Standalone()

	rootCmd.AddCommand(getDesktopForWindowCmd)

	carapace.Gen(getDesktopForWindowCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
