package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Import history for the current shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_autoCmd).Standalone()

	help_importCmd.AddCommand(help_import_autoCmd)
}
