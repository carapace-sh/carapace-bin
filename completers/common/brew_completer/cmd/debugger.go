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

	debuggerCmd.Flags().Bool("debug", false, "Display any debugging information.")
	debuggerCmd.Flags().Bool("help", false, "Show this message.")
	debuggerCmd.Flags().Bool("open", false, "Automatically open the debugger on failure.")
	debuggerCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	debuggerCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(debuggerCmd)
}
