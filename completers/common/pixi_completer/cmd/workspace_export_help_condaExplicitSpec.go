package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_export_help_condaExplicitSpecCmd = &cobra.Command{
	Use:   "conda-explicit-spec",
	Short: "Export workspace environment to a conda explicit specification file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_export_help_condaExplicitSpecCmd).Standalone()

	workspace_export_helpCmd.AddCommand(workspace_export_help_condaExplicitSpecCmd)
}
