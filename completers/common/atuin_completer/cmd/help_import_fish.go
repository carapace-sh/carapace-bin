package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_fishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Import history from the fish history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_fishCmd).Standalone()

	help_importCmd.AddCommand(help_import_fishCmd)
}
