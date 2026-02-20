package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Run an arbitrary process in a new window/tab",
}

func init() {
	launchCmd.AddCommand(atCmd)
	carapace.Gen(launchCmd).Standalone()

	launchCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(launchCmd).FlagCompletion(carapace.ActionMap{})
}