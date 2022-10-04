package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gum_completer/cmd/action"
	"github.com/spf13/cobra"
)

var inputCmd = &cobra.Command{
	Use:   "input",
	Short: "Prompt for some input",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputCmd).Standalone()

	inputCmd.Flags().String("char-limit", "", "Maximum value length")
	inputCmd.Flags().String("cursor.foreground", "", "Foreground Color")
	inputCmd.Flags().Bool("password", false, "Mask input characters")
	inputCmd.Flags().String("placeholder", "", "Placeholder value")
	inputCmd.Flags().String("prompt", "", "Prompt to display")
	inputCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	inputCmd.Flags().String("value", "", "Initial value")
	inputCmd.Flags().String("width", "", "Input width")
	rootCmd.AddCommand(inputCmd)

	carapace.Gen(inputCmd).FlagCompletion(carapace.ActionMap{
		"cursor.foreground": action.ActionColors(),
		"prompt.foreground": action.ActionColors(),
	})
}
