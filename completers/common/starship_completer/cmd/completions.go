package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:   "completions",
	Short: "Generate starship shell completions for your shell to stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	completionsCmd.Flags().BoolP("help", "h", false, "Prints help information")
	completionsCmd.Flags().BoolP("version", "V", false, "Prints version information")
	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "powershell", "elvish"),
	)
}
