package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Prints help for commands, or the index.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().String("help_verbosity", "", "Select the verbosity of the help command.")
	helpCmd.Flags().StringP("long", "l", "", "Show full description of each option, instead of just its name.")
	helpCmd.Flags().String("short", "", "Show only the names of the options, not their types or meanings.")
	rootCmd.AddCommand(helpCmd)
}
