package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var setDesktopForWindowCmd = &cobra.Command{
	Use:   "set_desktop_for_window",
	Short: "Move a window to a different desktop",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setDesktopForWindowCmd).Standalone()

	rootCmd.AddCommand(setDesktopForWindowCmd)

	carapace.Gen(setDesktopForWindowCmd).PositionalCompletion(
		xdotool.ActionWindows(),
		xdotool.ActionDesktops(),
	)
}
