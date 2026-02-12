package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_reshCmd = &cobra.Command{
	Use:   "resh",
	Short: "Import history from the resh history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_reshCmd).Standalone()

	import_reshCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_reshCmd)
}
