package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var clickCmd = &cobra.Command{
	Use:   "click",
	Short: "Send a click, that is a mousedown followed by mouseup for the given button",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clickCmd).Standalone()

	clickCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before clicking")
	clickCmd.Flags().String("delay", "", "Specify how long, in milliseconds, to delay between clicks")
	clickCmd.Flags().String("repeat", "", "Specify how many times to click")
	clickCmd.Flags().String("window", "", "Specify a window to send a click to")
	rootCmd.AddCommand(clickCmd)

	carapace.Gen(clickCmd).FlagCompletion(carapace.ActionMap{
		"window": xdotool.ActionWindows(),
	})

	carapace.Gen(clickCmd).PositionalCompletion(
		os.ActionMouseButtons(),
	)
}
