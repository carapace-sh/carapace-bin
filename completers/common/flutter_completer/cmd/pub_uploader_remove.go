package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_uploader_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove an email from the list of uploaders",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_uploader_removeCmd).Standalone()

	pub_uploader_removeCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_uploader_removeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_uploader_removeCmd.Flags().String("package", "", "The package whose uploaders will be modified.")
	pub_uploaderCmd.AddCommand(pub_uploader_removeCmd)
}
