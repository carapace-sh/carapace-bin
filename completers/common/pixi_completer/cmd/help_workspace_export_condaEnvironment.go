package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_workspace_export_condaEnvironmentCmd = &cobra.Command{
	Use:   "conda-environment",
	Short: "Export workspace environment to a conda environment.yaml file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_workspace_export_condaEnvironmentCmd).Standalone()

	help_workspace_exportCmd.AddCommand(help_workspace_export_condaEnvironmentCmd)
}
