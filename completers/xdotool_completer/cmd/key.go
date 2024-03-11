package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var keyCmd = &cobra.Command{
	Use:   "key",
	Short: "Type a given keystroke",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keyCmd).Standalone()

	keyCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before sending keystrokes")
	keyCmd.Flags().String("delay", "", "Delay between keystrokes")
	keyCmd.Flags().String("window", "", "Send keystrokes to a specific window id")
	rootCmd.AddCommand(keyCmd)

	carapace.Gen(keyCmd).FlagCompletion(carapace.ActionMap{
		"window": xdotool.ActionWindows(),
	})

	carapace.Gen(keyCmd).PositionalAnyCompletion(
		xdotool.ActionKeys().UniqueList("+"),
	)
}
