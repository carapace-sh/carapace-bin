package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gum_completer/cmd/action"
	"github.com/spf13/cobra"
)

var filterCmd = &cobra.Command{
	Use:   "filter",
	Short: "Filter items from a list",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(filterCmd).Standalone()

	filterCmd.Flags().String("height", "", "Input height")
	filterCmd.Flags().String("indicator", "", "Character for selection")
	filterCmd.Flags().String("indicator.foreground", "", "Foreground Color")
	filterCmd.Flags().String("limit", "", "Maximum number of options to pick")
	filterCmd.Flags().String("match.foreground", "", "Foreground Color")
	filterCmd.Flags().Bool("no-limit", false, "Pick unlimited number of options")
	filterCmd.Flags().String("placeholder", "", "Placeholder value")
	filterCmd.Flags().String("prompt", "", "Prompt to display")
	filterCmd.Flags().String("prompt.foreground", "", "Foreground Color")
	filterCmd.Flags().String("selected-indicator.foreground", "", "Foreground Color")
	filterCmd.Flags().String("selected-prefix", "", "Character to indicate selected items")
	filterCmd.Flags().String("text.foreground", "", "Foreground Color")
	filterCmd.Flags().String("unselected-prefix", "", "Character to indicate unselected items")
	filterCmd.Flags().String("unselected-prefix.foreground", "", "Foreground Color")
	filterCmd.Flags().String("value", "", "Initial filter value")
	filterCmd.Flags().String("width", "", "Input width")
	rootCmd.AddCommand(filterCmd)

	carapace.Gen(filterCmd).FlagCompletion(carapace.ActionMap{
		"indicator.foreground":          action.ActionColors(),
		"match.foreground":              action.ActionColors(),
		"prompt.foreground":             action.ActionColors(),
		"selected-indicator.foreground": action.ActionColors(),
		"text.foreground":               action.ActionColors(),
		"unselected-prefix.foreground":  action.ActionColors(),
	})
}
