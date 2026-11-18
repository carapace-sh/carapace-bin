package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateManCompletionsCmd = &cobra.Command{
	Use:     "generate-man-completions",
	Short:   "Generate Homebrew's manpages and shell completions",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateManCompletionsCmd).Standalone()

	generateManCompletionsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	generateManCompletionsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	generateManCompletionsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	generateManCompletionsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateManCompletionsCmd)
}
