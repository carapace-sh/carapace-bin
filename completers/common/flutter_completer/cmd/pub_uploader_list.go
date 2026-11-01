package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pub_uploader_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List the uploaders for a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pub_uploader_listCmd).Standalone()

	pub_uploader_listCmd.Flags().StringP("directory", "C", "", "Run this in the directory <dir>.")
	pub_uploader_listCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	pub_uploader_listCmd.Flags().String("package", "", "The package whose uploaders will be modified.")
	pub_uploaderCmd.AddCommand(pub_uploader_listCmd)
}
