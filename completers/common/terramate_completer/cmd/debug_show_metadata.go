package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_show_metadataCmd = &cobra.Command{
	Use:   "metadata",
	Short: "Shows metadata available on the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_show_metadataCmd).Standalone()

	debug_showCmd.AddCommand(debug_show_metadataCmd)
}
