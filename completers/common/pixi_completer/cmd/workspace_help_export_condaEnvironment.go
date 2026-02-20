package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspace_help_export_condaEnvironmentCmd = &cobra.Command{
	Use:   "conda-environment",
	Short: "Export workspace environment to a conda environment.yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspace_help_export_condaEnvironmentCmd).Standalone()

	workspace_help_exportCmd.AddCommand(workspace_help_export_condaEnvironmentCmd)
}
