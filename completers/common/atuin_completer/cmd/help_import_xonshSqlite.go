package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_xonshSqliteCmd = &cobra.Command{
	Use:   "xonsh-sqlite",
	Short: "Import history from xonsh sqlite db",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_xonshSqliteCmd).Standalone()

	help_importCmd.AddCommand(help_import_xonshSqliteCmd)
}
