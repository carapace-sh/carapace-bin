package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var import_powershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Import history from the powershell history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(import_powershellCmd).Standalone()

	import_powershellCmd.Flags().BoolP("help", "h", false, "Print help")
	importCmd.AddCommand(import_powershellCmd)
}
