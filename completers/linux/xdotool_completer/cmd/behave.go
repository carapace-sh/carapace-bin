package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var behaveCmd = &cobra.Command{
	Use:   "behave",
	Short: "Bind an action to an event on a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(behaveCmd).Standalone()

	rootCmd.AddCommand(behaveCmd)

	carapace.Gen(behaveCmd).PositionalCompletion(
		carapace.Batch(
			xdotool.ActionWindows(),
			carapace.ActionValues("%1", "%@"),
		).ToA(),
		carapace.ActionValuesDescribed(
			"blur", "Fires when the window loses focus",
			"focus", "Fires when the window gets input focus",
			"mouse-click", "Fires when the mouse is clicked",
			"mouse-enter", "Fires when the mouse enters a window",
			"mouse-leave", "Fires when the mouse leaves a window",
		),
	)
}
