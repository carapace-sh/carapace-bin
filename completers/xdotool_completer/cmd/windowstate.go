package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowstateCmd = &cobra.Command{
	Use:   "windowstate",
	Short: "Change a property on a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowstateCmd).Standalone()
	windowstateCmd.Flags().String("add", "", "add a property")
	windowstateCmd.Flags().String("remove", "", "remove a property")
	windowstateCmd.Flags().String("toggle", "", "toggle a property")

	rootCmd.AddCommand(windowstateCmd)

	carapace.Gen(windowstateCmd).FlagCompletion(carapace.ActionMap{
		"add":    xdotool.ActionProperties(),
		"remove": xdotool.ActionProperties(),
		"toggle": xdotool.ActionProperties(),
	})

	carapace.Gen(windowstateCmd).PositionalCompletion(
		xdotool.ActionWindows(),
	)
}
