package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Print this message or the help of the given subcommand(s)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	helpCmd.Flags().StringP("keyword", "k", "", "Show help for keywords instead of commands")
	rootCmd.AddCommand(helpCmd)

	helpCmd.MarkFlagsMutuallyExclusive("help", "keyword")

	carapace.Gen(helpCmd).FlagCompletion(carapace.ActionMap{
		"keyword": carapace.ActionValuesDescribed(
			"bookmarks", "Named pointers to revisions (similar to Git's branches)",
			"config", "How and where to set configuration options",
			"filesets", "A functional language for selecting a set of files",
			"glossary", "Definitions of various terms",
			"revsets", "A functional language for selecting a set of revision",
			"templates", "A functional language to customize command output",
			"tutorial", "Show a tutorial to get started with jj",
		),
	})

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(rootCmd),
	)
}
