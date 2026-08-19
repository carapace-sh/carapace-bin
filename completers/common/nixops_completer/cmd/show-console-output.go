package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ShowConsoleOutputCmd = &cobra.Command{
	Use:   "show-console-output",
	Short: "Show Console Output",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ShowConsoleOutputCmd).Standalone()
	rootCmd.AddCommand(ShowConsoleOutputCmd)
}
