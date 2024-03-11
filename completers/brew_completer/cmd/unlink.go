package cmd

import (
	"github.com/carapace-sh/carapace"
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

	unlinkCmd.Flags().Bool("debug", false, "Display any debugging information.")
	unlinkCmd.Flags().Bool("dry-run", false, "List files which would be unlinked without actually unlinking or deleting any files.")
	unlinkCmd.Flags().Bool("help", false, "Show this message.")
	unlinkCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	unlinkCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(unlinkCmd)
}
