package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var project_cleanupCmd = &cobra.Command{
	Use:   "cleanup",
	Short: "Clean resources associated with a project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(project_cleanupCmd).Standalone()

	project_cleanupCmd.Flags().Bool("auth-project", false, "Delete resources of the project used to authenticate")
	project_cleanupCmd.Flags().Bool("auto-approve", false, "Delete resources without asking for confirmation")
	project_cleanupCmd.Flags().String("created-before", "", "Only delete resources created before the given time")
	project_cleanupCmd.Flags().Bool("dry-run", false, "List a project's resources but do not delete them")
	project_cleanupCmd.Flags().String("project", "", "Project to clean (name or ID)")
	project_cleanupCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	project_cleanupCmd.Flags().String("skip-resource", "", "Skip cleanup of specific resource (repeat if necessary)")
	project_cleanupCmd.Flags().String("updated-before", "", "Only delete resources updated before the given time")
	projectCmd.AddCommand(project_cleanupCmd)
}
