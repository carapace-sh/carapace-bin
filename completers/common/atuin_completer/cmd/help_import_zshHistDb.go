package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_zshHistDbCmd = &cobra.Command{
	Use:   "zsh-hist-db",
	Short: "Import history from the zsh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_zshHistDbCmd).Standalone()

	help_importCmd.AddCommand(help_import_zshHistDbCmd)
}
