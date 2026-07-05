package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustSettingsExportCmd = &cobra.Command{
	Use:   "trust-settings-export",
	Short: "Export trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustSettingsExportCmd).Standalone()
	trustSettingsExportCmd.Flags().BoolP("admin", "d", false, "Export admin Trust Settings (default: user)")
	trustSettingsExportCmd.Flags().BoolP("system", "s", false, "Export system Trust Settings (default: user)")
	rootCmd.AddCommand(trustSettingsExportCmd)
	carapace.Gen(trustSettingsExportCmd).PositionalCompletion(carapace.ActionFiles())
}
