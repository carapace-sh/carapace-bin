package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dscv3_userSettingsFileCmd = &cobra.Command{
	Use:   "user-settings-file",
	Short: "Manage user settings file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dscv3_userSettingsFileCmd).Standalone()

	dscv3_userSettingsFileCmd.Flags().Bool("export", false, "Get all state instances")
	dscv3_userSettingsFileCmd.Flags().Bool("get", false, "Get the resource state")
	dscv3_userSettingsFileCmd.Flags().Bool("schema", false, "Get the resource schema")
	dscv3_userSettingsFileCmd.Flags().Bool("set", false, "Set the resource state")
	dscv3_userSettingsFileCmd.Flags().Bool("test", false, "Test the resource state")
	dscv3Cmd.AddCommand(dscv3_userSettingsFileCmd)
}
