package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var completionsCmd = &cobra.Command{
	Use:     "completions",
	Short:   "Generate shell completion scripts for bash, zsh, or fish",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionsCmd).Standalone()

	completionsCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(completionsCmd)

	carapace.Gen(completionsCmd).PositionalCompletion(
		carapace.ActionValues(
			"bash", "elvish", "fish", "powershell", "zsh",
		),
	)
}
