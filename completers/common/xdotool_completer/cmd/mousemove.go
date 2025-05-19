package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var mousemoveCmd = &cobra.Command{
	Use:   "mousemove",
	Short: "Move the mouse to the specific X and Y coordinates on the screen",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mousemoveCmd).Standalone()

	mousemoveCmd.Flags().Bool("clearmodifiers", false, "See CLEARMODIFIERS")
	mousemoveCmd.Flags().Bool("polar", false, "Use polar coordinates")
	mousemoveCmd.Flags().String("screen", "", "Move the mouse to the specified screen to move to")
	mousemoveCmd.Flags().Bool("sync", false, "After sending the mouse move request, wait until the mouse is actually moved")
	mousemoveCmd.Flags().String("window", "", "Specify a window to move relative to")
	rootCmd.AddCommand(mousemoveCmd)

	carapace.Gen(mousemoveCmd).FlagCompletion(carapace.ActionMap{
		"screen": os.ActionScreens(true),
		"window": xdotool.ActionWindows(),
	})
}
