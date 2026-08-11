package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset changes to a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().BoolP("help", "h", false, "Print help")
	resetCmd.Flags().String("last-merged-from", "", "If given, the files will be reset to the last point of merge from this branch, or the branch point from this branch if no merge has been performed")
	resetCmd.Flags().Bool("purge", false, "Delete untracked files")
	resetCmd.Flags().String("revision", "", "Revision to reset files to")
	resetCmd.Flags().String("targets", "", "Path to a targets file containing all the paths to all files")
	rootCmd.AddCommand(resetCmd)
}
