package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_fishCmd = &cobra.Command{
	Use:   "fish",
	Short: "Import history from the fish history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_fishCmd).Standalone()

	import_fishCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_fishCmd)
}
