package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:     "completions",
	Short:   "Control whether Homebrew automatically links external tap shell completion files",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	completionsCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	completionsCmd.Flags().BoolP("help", "h", false, "Show this message.")
	completionsCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	completionsCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).PositionalCompletion(
		carapace.ActionValues("link", "unlink"),
	)
}
