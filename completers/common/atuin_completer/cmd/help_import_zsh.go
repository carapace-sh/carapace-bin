package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Import history from the zsh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_zshCmd).Standalone()

	help_importCmd.AddCommand(help_import_zshCmd)
}
