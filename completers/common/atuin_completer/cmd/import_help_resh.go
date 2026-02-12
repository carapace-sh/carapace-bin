package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_help_reshCmd = &cobra.Command{
	Use:   "resh",
	Short: "Import history from the resh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_help_reshCmd).Standalone()

	import_helpCmd.AddCommand(import_help_reshCmd)
}
