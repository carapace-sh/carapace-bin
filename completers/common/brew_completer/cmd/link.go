package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/brew_completer/cmd/action"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Symlink all of <formula>'s installed files into Homebrew's prefix",
	GroupID: "main",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().Bool("HEAD", false, "Link the HEAD version of the formula if it is installed.")
	linkCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	linkCmd.Flags().BoolP("dry-run", "n", false, "List files which would be linked or deleted by `brew link --overwrite` without actually linking or deleting any files.")
	linkCmd.Flags().BoolP("force", "f", false, "Allow keg-only formulae to be linked.")
	linkCmd.Flags().BoolP("help", "h", false, "Show this message.")
	linkCmd.Flags().Bool("overwrite", false, "Delete files that already exist in the prefix while linking.")
	linkCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	linkCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalAnyCompletion(
		action.ActionList(linkCmd).FilterArgs(),
	)
}
