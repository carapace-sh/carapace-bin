package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unpinCmd = &cobra.Command{
	Use:     "unpin",
	Short:   "Unpin <formula>, allowing them to be upgraded by `brew upgrade` <formula>",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpinCmd).Standalone()

	unpinCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	unpinCmd.Flags().BoolP("help", "h", false, "Show this message.")
	unpinCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	unpinCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(unpinCmd)

	carapace.Gen(unpinCmd).PositionalAnyCompletion(
		action.ActionList(unpinCmd).FilterArgs(),
	)
}
