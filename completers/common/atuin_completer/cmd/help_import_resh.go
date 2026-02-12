package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_reshCmd = &cobra.Command{
	Use:   "resh",
	Short: "Import history from the resh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_reshCmd).Standalone()

	help_importCmd.AddCommand(help_import_reshCmd)
}
