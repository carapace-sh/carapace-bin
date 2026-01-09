package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dscv3_adminSettingsCmd = &cobra.Command{
	Use:   "admin-settings",
	Short: "Manage administrator settings",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dscv3_adminSettingsCmd).Standalone()

	dscv3_adminSettingsCmd.Flags().Bool("export", false, "Get all state instances")
	dscv3_adminSettingsCmd.Flags().Bool("get", false, "Get the resource state")
	dscv3_adminSettingsCmd.Flags().Bool("schema", false, "Get the resource schema")
	dscv3_adminSettingsCmd.Flags().Bool("set", false, "Set the resource state")
	dscv3_adminSettingsCmd.Flags().Bool("test", false, "Test the resource state")
	dscv3Cmd.AddCommand(dscv3_adminSettingsCmd)
}
