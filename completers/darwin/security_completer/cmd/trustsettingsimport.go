package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var trustSettingsImportCmd = &cobra.Command{
	Use:   "trust-settings-import",
	Short: "Import trust settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trustSettingsImportCmd).Standalone()
	trustSettingsImportCmd.Flags().BoolP("admin", "d", false, "Import admin Trust Settings (default: user)")
	rootCmd.AddCommand(trustSettingsImportCmd)
	carapace.Gen(trustSettingsImportCmd).PositionalCompletion(carapace.ActionFiles())
}
