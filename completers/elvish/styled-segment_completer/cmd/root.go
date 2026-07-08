package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "styled-segment",
	Short: "Create a styled text segment",
	Long:  "https://elv.sh/ref/builtin.html#styled-segment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("bg-color", "bg-color", "default", "background color")
	rootCmd.Flags().BoolS("blink", "blink", false, "blink text")
	rootCmd.Flags().BoolS("bold", "bold", false, "bold text")
	rootCmd.Flags().BoolS("dim", "dim", false, "dim text")
	rootCmd.Flags().StringS("fg-color", "fg-color", "default", "foreground color")
	rootCmd.Flags().BoolS("inverse", "inverse", false, "inverse text")
	rootCmd.Flags().BoolS("italic", "italic", false, "italic text")
	rootCmd.Flags().BoolS("underlined", "underlined", false, "underlined text")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"bg-color": carapace.ActionValuesDescribed(
			"default", "default color",
			"black", "black",
			"red", "red",
			"green", "green",
			"yellow", "yellow",
			"blue", "blue",
			"magenta", "magenta",
			"cyan", "cyan",
			"white", "white",
			"bright-black", "bright black",
			"bright-red", "bright red",
			"bright-green", "bright green",
			"bright-yellow", "bright yellow",
			"bright-blue", "bright blue",
			"bright-magenta", "bright magenta",
			"bright-cyan", "bright cyan",
			"bright-white", "bright white",
		).Usage("background color"),
		"fg-color": carapace.ActionValuesDescribed(
			"default", "default color",
			"black", "black",
			"red", "red",
			"green", "green",
			"yellow", "yellow",
			"blue", "blue",
			"magenta", "magenta",
			"cyan", "cyan",
			"white", "white",
			"bright-black", "bright black",
			"bright-red", "bright red",
			"bright-green", "bright green",
			"bright-yellow", "bright yellow",
			"bright-blue", "bright blue",
			"bright-magenta", "bright magenta",
			"bright-cyan", "bright cyan",
			"bright-white", "bright white",
		).Usage("foreground color"),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("object to style"),
	)
}
