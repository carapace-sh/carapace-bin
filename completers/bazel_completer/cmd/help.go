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

	helpCmd.Flags().Bool("help_verbosity", false, "Select the verbosity of the help command.")
	helpCmd.Flags().BoolP("long", "l", false, "Show full description of each option, instead of just its name.")
	helpCmd.Flags().Bool("short", false, "Show only the names of the options, not their types or meanings.")
	rootCmd.AddCommand(helpCmd)
}
