package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var extension_browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Enter a UI for browsing, adding, and removing extensions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_browseCmd).Standalone()

	extension_browseCmd.Flags().Bool("debug", false, "log to /tmp/extBrowse-*")
	extensionCmd.AddCommand(extension_browseCmd)
}
