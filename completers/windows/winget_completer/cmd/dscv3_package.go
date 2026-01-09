package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dscv3_packageCmd = &cobra.Command{
	Use:   "package",
	Short: "Manage package state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dscv3_packageCmd).Standalone()

	dscv3_packageCmd.Flags().Bool("export", false, "Get all state instances")
	dscv3_packageCmd.Flags().Bool("get", false, "Get the resource state")
	dscv3_packageCmd.Flags().Bool("schema", false, "Get the resource schema")
	dscv3_packageCmd.Flags().Bool("set", false, "Set the resource state")
	dscv3_packageCmd.Flags().Bool("test", false, "Test the resource state")
	dscv3Cmd.AddCommand(dscv3_packageCmd)
}
