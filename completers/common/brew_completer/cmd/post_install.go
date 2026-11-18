package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var post_installCmd = &cobra.Command{
	Use:   "post_install",
	Short: "Rerun the post-install steps for <formula>",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(post_installCmd).Standalone()

	post_installCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	post_installCmd.Flags().BoolP("help", "h", false, "Show this message.")
	post_installCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	post_installCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(post_installCmd)

	carapace.Gen(post_installCmd).PositionalAnyCompletion(
		action.ActionList(post_installCmd).FilterArgs(),
	)
}
