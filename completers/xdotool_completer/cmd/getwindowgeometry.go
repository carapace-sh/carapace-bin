package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var getwindowgeometryCmd = &cobra.Command{
	Use:   "getwindowgeometry",
	Short: "Output the geometry (location and position) of a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowgeometryCmd).Standalone()
	getwindowgeometryCmd.Flags().Bool("shell", false, "Output values suitable for 'eval' in a shell")

	rootCmd.AddCommand(getwindowgeometryCmd)

	carapace.Gen(getwindowgeometryCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
