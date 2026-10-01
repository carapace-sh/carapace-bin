package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generate shell completions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(completionCmd).Standalone()

	completionCmd.Flags().Bool("force", false, "Replace a file at the target path that mise did not write")
	completionCmd.Flags().Bool("include-bash-completion-lib", false, "Retained for compatibility with older generators")
	completionCmd.Flags().Bool("install", false, "Install the script where this shell looks for it")
	completionCmd.Flags().String("tool", "", "A tool's completion instead of mise's own")
	rootCmd.AddCommand(completionCmd)

	carapace.Gen(completionCmd).FlagCompletion(carapace.ActionMap{
		"tool": action.ActionTools(),
	})

	carapace.Gen(completionCmd).PositionalCompletion(
		action.ActionShells(),
	)
}
