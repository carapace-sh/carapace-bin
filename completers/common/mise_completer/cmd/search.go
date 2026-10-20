package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for available tools",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(searchCmd).Standalone()

	searchCmd.Flags().BoolP("all", "a", false, "Search every source: the registry, aqua, installed backend plugins, and the npm, cargo, gem, and dotnet package registries")
	searchCmd.Flags().Bool("complete", false, "Print all tools with descriptions for shell completions")
	searchCmd.Flags().Bool("complete-ids", false, "Print only tool identifiers for shell completions")
	searchCmd.Flags().BoolP("help", "h", false, "Print help")
	searchCmd.Flags().BoolP("interactive", "i", false, "Show an interactive search menu")
	searchCmd.Flags().StringP("match-type", "m", "fuzzy", "Match type: equal, contains, or fuzzy")
	searchCmd.Flags().Bool("no-header", false, "Do not print the table header")
	searchCmd.Flags().Bool("no-headers", false, "Do not print the table header")
	searchCmd.Flag("complete").Hidden = true
	searchCmd.Flag("complete-ids").Hidden = true
	searchCmd.Flag("no-headers").Hidden = true
	rootCmd.AddCommand(searchCmd)

	carapace.Gen(searchCmd).FlagCompletion(carapace.ActionMap{
		"match-type": carapace.ActionValues("equal", "contains", "fuzzy"),
	})
}
