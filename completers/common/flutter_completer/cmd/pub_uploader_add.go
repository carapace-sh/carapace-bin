package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_uploader_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add an email to the list of uploaders",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_uploader_addCmd).Standalone()

	pub_uploader_addCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_uploader_addCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_uploader_addCmd.Flags().String("package", "", "The package whose uploaders will be modified.")
	pub_uploaderCmd.AddCommand(pub_uploader_addCmd)
}
