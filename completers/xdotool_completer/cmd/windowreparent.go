package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var windowreparentCmd = &cobra.Command{
	Use:   "windowreparent",
	Short: "Reparent a window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(windowreparentCmd).Standalone()

	rootCmd.AddCommand(windowreparentCmd)

	carapace.Gen(windowreparentCmd).PositionalAnyCompletion(
		xdotool.ActionWindows().FilterArgs(), // TODO verify
	)
}
