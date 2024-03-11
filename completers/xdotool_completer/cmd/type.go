package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/xdotool"
	"github.com/spf13/cobra"
)

var typeCmd = &cobra.Command{
	Use:   "type",
	Short: "Type as if you had typed it",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(typeCmd).Standalone()

	typeCmd.Flags().Bool("clearmodifiers", false, "Clear modifiers before sending keystrokes")
	typeCmd.Flags().String("delay", "", "Delay between keystrokes")
	typeCmd.Flags().String("window", "", "Send keystrokes to a specific window id")
	rootCmd.AddCommand(typeCmd)

	carapace.Gen(typeCmd).FlagCompletion(carapace.ActionMap{
		"window": xdotool.ActionWindows(),
	})
}
