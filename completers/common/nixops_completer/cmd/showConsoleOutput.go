package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showConsoleOutputCmd = &cobra.Command{
	Use:   "show-console-output",
	Short: "print the machine console output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showConsoleOutputCmd).Standalone()
	rootCmd.AddCommand(showConsoleOutputCmd)
}
