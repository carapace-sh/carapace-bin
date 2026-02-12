package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_xonshCmd = &cobra.Command{
	Use:   "xonsh",
	Short: "Import history from xonsh json files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_xonshCmd).Standalone()

	help_importCmd.AddCommand(help_import_xonshCmd)
}
