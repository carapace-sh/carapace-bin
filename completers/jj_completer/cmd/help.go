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

	helpCmd.Flags().Bool("help", false, "Print help (see more with '--help')")
	helpCmd.Flags().StringArrayP("keyword", "k", nil, "Show help for keywords instead of commands")
	rootCmd.AddCommand(helpCmd)

	helpCmd.MarkFlagsMutuallyExclusive("help", "keyword")

	carapace.Gen(helpCmd).FlagCompletion(carapace.ActionMap{
		"keyword": carapace.ActionValues("revsets", "tutorial"),
	})

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(rootCmd),
	)
}
