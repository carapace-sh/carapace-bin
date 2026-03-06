package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hook_postCreateCmd = &cobra.Command{
	Use:   "post-create",
	Short: "Run post-create hooks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hook_postCreateCmd).Standalone()

	hook_postCreateCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hook_postCreateCmd.Flags().StringSlice("var", nil, "Override built-in template variable (KEY=VALUE)")
	hook_postCreateCmd.Flags().BoolP("yes", "y", false, "Skip approval prompts")
	hookCmd.AddCommand(hook_postCreateCmd)

	// TODO complete names
}
