package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_preMergeCmd = &cobra.Command{
	Use:   "pre-merge",
	Short: "Run pre-merge hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_preMergeCmd).Standalone()

	hook_preMergeCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_preMergeCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_preMergeCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_preMergeCmd)

	// TODO complete names
}
