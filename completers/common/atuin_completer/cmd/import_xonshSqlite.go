package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_xonshSqliteCmd = &cobra.Command{
	Use:   "xonsh-sqlite",
	Short: "Import history from xonsh sqlite db",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_xonshSqliteCmd).Standalone()

	import_xonshSqliteCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_xonshSqliteCmd)
}
