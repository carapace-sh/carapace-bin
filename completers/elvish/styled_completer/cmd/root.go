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

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("object to style"),
	)

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValuesDescribed(
			"bold", "bold text",
			"no-bold", "remove bold",
			"toggle-bold", "toggle bold",
			"dim", "dim text",
			"no-dim", "remove dim",
			"toggle-dim", "toggle dim",
			"italic", "italic text",
			"no-italic", "remove italic",
			"toggle-italic", "toggle italic",
			"underlined", "underlined text",
			"no-underlined", "remove underlined",
			"toggle-underlined", "toggle underlined",
			"blink", "blinking text",
			"no-blink", "remove blink",
			"toggle-blink", "toggle blink",
			"inverse", "inverse text",
			"no-inverse", "remove inverse",
			"toggle-inverse", "toggle inverse",
			"black", "black foreground",
			"red", "red foreground",
			"green", "green foreground",
			"yellow", "yellow foreground",
			"blue", "blue foreground",
			"magenta", "magenta foreground",
			"cyan", "cyan foreground",
			"white", "white foreground",
			"bright-black", "bright black foreground",
			"bright-red", "bright red foreground",
			"bright-green", "bright green foreground",
			"bright-yellow", "bright yellow foreground",
			"bright-blue", "bright blue foreground",
			"bright-magenta", "bright magenta foreground",
			"bright-cyan", "bright cyan foreground",
			"bright-white", "bright white foreground",
			"bg-black", "black background",
			"bg-red", "red background",
			"bg-green", "green background",
			"bg-yellow", "yellow background",
			"bg-blue", "blue background",
			"bg-magenta", "magenta background",
			"bg-cyan", "cyan background",
			"bg-white", "white background",
			"bg-bright-black", "bright black background",
			"bg-bright-red", "bright red background",
			"bg-bright-green", "bright green background",
			"bg-bright-yellow", "bright yellow background",
			"bg-bright-blue", "bright blue background",
			"bg-bright-magenta", "bright magenta background",
			"bg-bright-cyan", "bright cyan background",
			"bg-bright-white", "bright white background",
		).Usage("style transformer"),
	)
}
