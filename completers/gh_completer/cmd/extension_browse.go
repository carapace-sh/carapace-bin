package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var extension_browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Enter a UI for browsing, adding, and removing extensions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_browseCmd).Standalone()

	extension_browseCmd.Flags().Bool("debug", false, "Log to /tmp/extBrowse-*")
	extension_browseCmd.Flags().BoolP("single-column", "s", false, "Render TUI with only one column of text")
	extensionCmd.AddCommand(extension_browseCmd)
}
