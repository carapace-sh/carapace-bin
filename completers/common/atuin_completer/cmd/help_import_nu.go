package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_nuCmd = &cobra.Command{
	Use:   "nu",
	Short: "Import history from the nu history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_nuCmd).Standalone()

	help_importCmd.AddCommand(help_import_nuCmd)
}
