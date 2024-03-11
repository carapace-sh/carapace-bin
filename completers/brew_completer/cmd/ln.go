package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lnCmd = &cobra.Command{
	Use:   "ln",
	Short: "Symlink all of <formula>'s installed files into Homebrew's prefix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lnCmd).Standalone()

	lnCmd.Flags().Bool("HEAD", false, "Link the HEAD version of the formula if it is installed.")
	lnCmd.Flags().Bool("debug", false, "Display any debugging information.")
	lnCmd.Flags().Bool("dry-run", false, "List files which would be linked or deleted by `brew link --overwrite` without actually linking or deleting any files.")
	lnCmd.Flags().Bool("force", false, "Allow keg-only formulae to be linked.")
	lnCmd.Flags().Bool("help", false, "Show this message.")
	lnCmd.Flags().Bool("overwrite", false, "Delete files that already exist in the prefix while linking.")
	lnCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	lnCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(lnCmd)
}
