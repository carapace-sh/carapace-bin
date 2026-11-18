package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var execCmd = &cobra.Command{
	Use:     "exec",
	Aliases: []string{"x"},
	Short:   "Run <command> in an environment populated by Homebrew formulae",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(execCmd).Standalone()

	execCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	execCmd.Flags().BoolP("help", "h", false, "Show this message.")
	execCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	execCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(execCmd)

	carapace.Gen(execCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
