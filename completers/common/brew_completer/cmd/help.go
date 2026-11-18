package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:     "help",
	Short:   "Output the usage instructions for `brew` <command>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpCmd).Standalone()

	helpCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	helpCmd.Flags().BoolP("help", "h", false, "Show this message.")
	helpCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	helpCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(helpCmd)

	carapace.Gen(helpCmd).PositionalAnyCompletion(
		carapace.ActionCommands(rootCmd),
	)
}
