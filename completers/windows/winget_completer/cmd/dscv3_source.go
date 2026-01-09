package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dscv3_sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage source configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dscv3_sourceCmd).Standalone()

	dscv3_sourceCmd.Flags().Bool("export", false, "Get all state instances")
	dscv3_sourceCmd.Flags().Bool("get", false, "Get the resource state")
	dscv3_sourceCmd.Flags().Bool("schema", false, "Get the resource schema")
	dscv3_sourceCmd.Flags().Bool("set", false, "Set the resource state")
	dscv3_sourceCmd.Flags().Bool("test", false, "Test the resource state")
	dscv3Cmd.AddCommand(dscv3_sourceCmd)
}
