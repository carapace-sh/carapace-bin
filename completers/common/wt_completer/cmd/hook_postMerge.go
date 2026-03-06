package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_postMergeCmd = &cobra.Command{
	Use:   "post-merge",
	Short: "Run post-merge hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_postMergeCmd).Standalone()

	hook_postMergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_postMergeCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_postMergeCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_postMergeCmd)

	// TODO complete names
}
