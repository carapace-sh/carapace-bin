package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_export_condaExplicitSpecCmd = &cobra.Command{
	Use:   "conda-explicit-spec",
	Short: "Export workspace environment to a conda explicit specification file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_export_condaExplicitSpecCmd).Standalone()

	help_workspace_exportCmd.AddCommand(help_workspace_export_condaExplicitSpecCmd)
}
