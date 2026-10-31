package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateShellCompletionCmd = &cobra.Command{
	Use:    "generate-shell-completion",
	Short:  "Generate shell completion",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateShellCompletionCmd).Standalone()

	generateShellCompletionCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(generateShellCompletionCmd)
}
