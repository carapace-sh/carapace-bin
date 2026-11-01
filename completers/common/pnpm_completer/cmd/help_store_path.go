package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_store_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Returns the path to the active store directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_store_pathCmd).Standalone()

	help_storeCmd.AddCommand(help_store_pathCmd)
}
