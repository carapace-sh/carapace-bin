package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var unpackCmd = &cobra.Command{
	Use:     "unpack",
	Short:   "Unpack the source files for <formula> into subdirectories of the current working directory",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unpackCmd).Standalone()

	unpackCmd.Flags().Bool("debug", false, "Display any debugging information.")
	unpackCmd.Flags().Bool("destdir", false, "Create subdirectories in the directory named by <path> instead.")
	unpackCmd.Flags().Bool("force", false, "Overwrite the destination directory if it already exists.")
	unpackCmd.Flags().Bool("git", false, "Initialise a Git repository in the unpacked source. This is useful for creating patches for the software.")
	unpackCmd.Flags().Bool("help", false, "Show this message.")
	unpackCmd.Flags().Bool("patch", false, "Patches for <formula> will be applied to the unpacked source.")
	unpackCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	unpackCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(unpackCmd)
}
