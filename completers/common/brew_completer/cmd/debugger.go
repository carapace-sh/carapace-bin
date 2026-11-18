package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debuggerCmd = &cobra.Command{
	Use:     "debugger",
	Short:   "Run the specified Homebrew command in debug mode",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debuggerCmd).Standalone()

	debuggerCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	debuggerCmd.Flags().BoolP("help", "h", false, "Show this message.")
	debuggerCmd.Flags().BoolP("open", "O", false, "Automatically open the debugger on failure.")
	debuggerCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	debuggerCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(debuggerCmd)

	carapace.Gen(debuggerCmd).PositionalAnyCompletion(
		carapace.ActionCommands(rootCmd),
	)
}
