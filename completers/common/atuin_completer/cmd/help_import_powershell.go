package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_import_powershellCmd = &cobra.Command{
	Use:   "powershell",
	Short: "Import history from the powershell history file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_import_powershellCmd).Standalone()

	help_importCmd.AddCommand(help_import_powershellCmd)
}
