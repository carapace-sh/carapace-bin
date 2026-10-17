package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var presetCmd = &cobra.Command{
	Use:   "preset",
	Short: "Prints a preset config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(presetCmd).Standalone()

	presetCmd.Flags().BoolP("force", "f", false, "Forcibly overwrite the output file if it already exists")
	presetCmd.Flags().BoolP("help", "h", false, "Print help")
	presetCmd.Flags().BoolP("list", "l", false, "List out all preset names")
	presetCmd.Flags().StringP("output", "o", "", "Output the preset to a file instead of stdout")
	rootCmd.AddCommand(presetCmd)

	carapace.Gen(presetCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})

	carapace.Gen(presetCmd).PositionalCompletion(
		carapace.ActionValues(
			"bracketed-segments",
			"catppuccin-powerline",
			"gruvbox-rainbow",
			"jetpack",
			"nerd-font-symbols",
			"no-empty-icons",
			"no-nerd-font",
			"no-runtime-versions",
			"pastel-powerline",
			"plain-text-symbols",
			"pure-preset",
			"tokyo-night",
		),
	)
}
