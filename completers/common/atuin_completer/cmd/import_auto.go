package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_autoCmd = &cobra.Command{
	Use:   "auto",
	Short: "Import history for the current shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_autoCmd).Standalone()

	import_autoCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_autoCmd)
}
