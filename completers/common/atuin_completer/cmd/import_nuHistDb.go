package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_nuHistDbCmd = &cobra.Command{
	Use:   "nu-hist-db",
	Short: "Import history from the nu history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_nuHistDbCmd).Standalone()

	import_nuHistDbCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_nuHistDbCmd)
}
