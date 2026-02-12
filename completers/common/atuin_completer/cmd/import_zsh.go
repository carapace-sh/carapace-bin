package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_zshCmd = &cobra.Command{
	Use:   "zsh",
	Short: "Import history from the zsh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_zshCmd).Standalone()

	import_zshCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_zshCmd)
}
