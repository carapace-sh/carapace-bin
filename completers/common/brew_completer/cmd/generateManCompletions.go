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

	generateManCompletionsCmd.Flags().Bool("debug", false, "Display any debugging information.")
	generateManCompletionsCmd.Flags().Bool("help", false, "Show this message.")
	generateManCompletionsCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	generateManCompletionsCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateManCompletionsCmd)
}
