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

	util_completionCmd.Flags().Bool("bash", false, "Print a completion script for Bash")
	util_completionCmd.Flags().Bool("fish", false, "Print a completion script for Fish")
	util_completionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	util_completionCmd.Flags().Bool("zsh", false, "Print a completion script for Zsh")
	utilCmd.AddCommand(util_completionCmd)
}
