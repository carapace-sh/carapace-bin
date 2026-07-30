package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ExportImageCmd = &cobra.Command{
	Use:   "Export-Image",
	Short: "export an image to a new file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ExportImageCmd).Standalone()
	rootCmd.AddCommand(ExportImageCmd)
}
