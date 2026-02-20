package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unicodeInputCmd = &cobra.Command{
	Use:   "unicode-input",
	Short: "Browse and select unicode characters by name",
}

func init() {
	rootCmd.AddCommand(unicodeInputCmd)
	carapace.Gen(unicodeInputCmd).Standalone()

	unicodeInputCmd.Flags().String("emoji-variation", "", "Whether to use the textual or the graphical form for emoji. By default the default form specified in the Unicode standard for the symbol is used.")
	unicodeInputCmd.Flags().BoolP("help", "h", false, "Show help for this command")
	unicodeInputCmd.Flags().String("tab", "", "The initial tab to display. Defaults to using the tab from the previous kitten invocation.")

	carapace.Gen(unicodeInputCmd).FlagCompletion(carapace.ActionMap{
		"emoji-variation": carapace.ActionValues("none", "graphic", "text"),
		"tab":             carapace.ActionValues("previous", "code", "emoticons", "favorites", "name"),
	})
}
