package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var util_completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Print a command-line-completion script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(util_completionCmd).Standalone()

	util_completionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	utilCmd.AddCommand(util_completionCmd)
}
