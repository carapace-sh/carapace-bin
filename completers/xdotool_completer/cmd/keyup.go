package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var keyupCmd = &cobra.Command{
	Use:   "keyup",
	Short: "Send keyup event for a given keystroke",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyupCmd).Standalone()

	keyupCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before sending keystrokes")
	keyupCmd.Flags().String("delay", "", "Delay between keystrokes")
	keyupCmd.Flags().String("window", "", "Send keystrokes to a specific window id")
	rootCmd.AddCommand(keyupCmd)

	carapace.Gen(keyupCmd).FlagCompletion(carapace.ActionMap{
		"window": xdotool.ActionWindows(),
	})

	carapace.Gen(keyupCmd).PositionalAnyCompletion(
		xdotool.ActionKeys().UniqueList("+"),
	)
}
