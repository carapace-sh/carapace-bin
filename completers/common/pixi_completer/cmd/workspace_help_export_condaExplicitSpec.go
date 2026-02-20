package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_export_condaExplicitSpecCmd = &cobra.Command{
	Use:   "conda-explicit-spec",
	Short: "Export workspace environment to a conda explicit specification file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_export_condaExplicitSpecCmd).Standalone()

	workspace_help_exportCmd.AddCommand(workspace_help_export_condaExplicitSpecCmd)
}
