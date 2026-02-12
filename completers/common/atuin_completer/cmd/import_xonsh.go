package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_xonshCmd = &cobra.Command{
	Use:   "xonsh",
	Short: "Import history from xonsh json files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_xonshCmd).Standalone()

	import_xonshCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_xonshCmd)
}
