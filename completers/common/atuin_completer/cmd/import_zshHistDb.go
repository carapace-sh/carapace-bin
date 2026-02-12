package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_zshHistDbCmd = &cobra.Command{
	Use:   "zsh-hist-db",
	Short: "Import history from the zsh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_zshHistDbCmd).Standalone()

	import_zshHistDbCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_zshHistDbCmd)
}
