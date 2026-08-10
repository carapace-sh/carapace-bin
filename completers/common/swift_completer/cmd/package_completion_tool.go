package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_completionToolCmd = &cobra.Command{
	Use:   "completion-tool",
	Short: "Command to generate shell completions",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_completionToolCmd).Standalone()
	package_completionToolCmd.Flags().SetInterspersed(false)

	package_completionToolCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_completionToolCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_completionToolCmd)

	carapace.Gen(package_completionToolCmd).PositionalCompletion(
		carapace.ActionValues(
			"generate-bash-script",
			"generate-zsh-script",
			"generate-fish-script",
			"list-dependencies",
			"list-executables",
			"list-snippets",
		),
	)
}
