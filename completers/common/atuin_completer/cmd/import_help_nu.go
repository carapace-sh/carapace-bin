package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_help_nuCmd = &cobra.Command{
	Use:   "nu",
	Short: "Import history from the nu history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_help_nuCmd).Standalone()

	import_helpCmd.AddCommand(import_help_nuCmd)
}
