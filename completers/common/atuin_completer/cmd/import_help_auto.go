package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_help_autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Import history for the current shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_help_autoCmd).Standalone()

	import_helpCmd.AddCommand(import_help_autoCmd)
}
