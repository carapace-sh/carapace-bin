package cmd

import (
	"github.com/carapace-sh/carapace"
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
	linkCmd.Flags().Bool("debug", false, "Display any debugging information.")
	linkCmd.Flags().Bool("dry-run", false, "List files which would be linked or deleted by `brew link --overwrite` without actually linking or deleting any files.")
	linkCmd.Flags().Bool("force", false, "Allow keg-only formulae to be linked.")
	linkCmd.Flags().Bool("help", false, "Show this message.")
	linkCmd.Flags().Bool("overwrite", false, "Delete files that already exist in the prefix while linking.")
	linkCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	linkCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(linkCmd)
}
