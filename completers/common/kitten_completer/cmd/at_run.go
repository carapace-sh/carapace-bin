package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a program on the computer in which kitty is running and get the output",
}

func init() {
	runCmd.AddCommand(atCmd)
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().BoolP("help", "h", false, "Show help for this command")

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{})
}