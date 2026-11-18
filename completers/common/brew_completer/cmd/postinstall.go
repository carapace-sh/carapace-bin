package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var postinstallCmd = &cobra.Command{
	Use:     "postinstall",
	Short:   "Rerun the post-install steps for <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(postinstallCmd).Standalone()

	postinstallCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	postinstallCmd.Flags().BoolP("help", "h", false, "Show this message.")
	postinstallCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	postinstallCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(postinstallCmd)

	carapace.Gen(postinstallCmd).PositionalAnyCompletion(
		action.ActionList(postinstallCmd).FilterArgs(),
	)
}
