package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var plugin_import_fromCmd = &cobra.Command{
	Use:   "from",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(plugin_import_fromCmd).Standalone()

	plugin_importCmd.AddCommand(plugin_import_fromCmd)
}
