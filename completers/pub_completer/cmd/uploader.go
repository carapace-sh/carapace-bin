package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var uploaderCmd = &cobra.Command{
	Use:   "uploader",
	Short: "Manage uploaders for a package on pub.dartlang.org",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(uploaderCmd).Standalone()

	uploaderCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	uploaderCmd.Flags().String("package", "", "The package whose uploaders will be modified.")
	rootCmd.AddCommand(uploaderCmd)

	carapace.Gen(uploaderCmd).PositionalCompletion(
		carapace.ActionValues("add", "remove"),
	)
}
