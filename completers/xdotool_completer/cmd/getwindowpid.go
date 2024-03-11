package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var getwindowpidCmd = &cobra.Command{
	Use:   "getwindowpid",
	Short: "Output the PID owning a given window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getwindowpidCmd).Standalone()

	rootCmd.AddCommand(getwindowpidCmd)

	carapace.Gen(getwindowpidCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
