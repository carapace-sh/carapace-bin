package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help [command]",
	Short: "Help about any command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionCommands(rootCmd),
			carapace.ActionValuesDescribed(
				"accessibility", "Learn about GitHub CLI's accessibility experiences",
				"actions", "Learn about working with GitHub Actions",
				"environment", "Environment variables that can be used with gh",
				"exit-codes", "Exit codes used by gh",
				"formatting", "Formatting options for JSON data exported from gh",
				"mintty", "Information about using gh with MinTTY",
				"reference", "A comprehensive reference of all gh commands",
			).Tag("help topics"),
		).ToA(),
	)

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(rootCmd),
	)
}
