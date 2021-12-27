package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var project_inspectCmd = &cobra.Command{
	Use:   "inspect",
	Short: "Inspect the details of a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_inspectCmd).Standalone()

	project_inspectCmd.Flags().Bool("json", false, "Output project information as JSON.")

	addGlobalOptions(project_inspectCmd)

	projectCmd.AddCommand(project_inspectCmd)
}
