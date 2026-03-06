package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_preCommitCmd = &cobra.Command{
	Use:   "pre-commit",
	Short: "Run pre-commit hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_preCommitCmd).Standalone()

	hook_preCommitCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_preCommitCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_preCommitCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_preCommitCmd)

	// TODO complete names
}
