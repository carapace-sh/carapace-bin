package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "styled",
	Short: "Style text",
	Long:  "https://elv.sh/ref/builtin.html#styled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetPrefix('&')

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("object to style"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"bg-black", "black background",
			"bg-blue", "blue background",
			"bg-bright-black", "bright black background",
			"bg-bright-blue", "bright blue background",
			"bg-bright-cyan", "bright cyan background",
			"bg-bright-green", "bright green background",
			"bg-bright-magenta", "bright magenta background",
			"bg-bright-red", "bright red background",
			"bg-bright-white", "bright white background",
			"bg-bright-yellow", "bright yellow background",
			"bg-cyan", "cyan background",
			"bg-green", "green background",
			"bg-magenta", "magenta background",
			"bg-red", "red background",
			"bg-white", "white background",
			"bg-yellow", "yellow background",
			"black", "black foreground",
			"blink", "blinking text",
			"blue", "blue foreground",
			"bold", "bold text",
			"bright-black", "bright black foreground",
			"bright-blue", "bright blue foreground",
			"bright-cyan", "bright cyan foreground",
			"bright-green", "bright green foreground",
			"bright-magenta", "bright magenta foreground",
			"bright-red", "bright red foreground",
			"bright-white", "bright white foreground",
			"bright-yellow", "bright yellow foreground",
			"cyan", "cyan foreground",
			"dim", "dim text",
			"green", "green foreground",
			"inverse", "inverse text",
			"italic", "italic text",
			"magenta", "magenta foreground",
			"no-blink", "remove blink",
			"no-bold", "remove bold",
			"no-dim", "remove dim",
			"no-inverse", "remove inverse",
			"no-italic", "remove italic",
			"no-underlined", "remove underlined",
			"red", "red foreground",
			"toggle-blink", "toggle blink",
			"toggle-bold", "toggle bold",
			"toggle-dim", "toggle dim",
			"toggle-inverse", "toggle inverse",
			"toggle-italic", "toggle italic",
			"toggle-underlined", "toggle underlined",
			"underlined", "underlined text",
			"white", "white foreground",
			"yellow", "yellow foreground",
		).Usage("style transformer"),
	)
}
