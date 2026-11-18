package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unlinkCmd = &cobra.Command{
	Use:     "unlink",
	Short:   "Remove symlinks for <formula> from Homebrew's prefix",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unlinkCmd).Standalone()

	unlinkCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	unlinkCmd.Flags().BoolP("dry-run", "n", false, "List files which would be unlinked without actually unlinking or deleting any files.")
	unlinkCmd.Flags().BoolP("help", "h", false, "Show this message.")
	unlinkCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	unlinkCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(unlinkCmd)

	carapace.Gen(unlinkCmd).PositionalAnyCompletion(
		action.ActionList(unlinkCmd).FilterArgs(),
	)
}
