package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var settings_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(settings_exportCmd).Standalone()

	settingsCmd.AddCommand(settings_exportCmd)
}
