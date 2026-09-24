package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_remove_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Dissociate a project from an endpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_remove_projectCmd).Standalone()

	endpoint_remove_projectCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	endpoint_removeCmd.AddCommand(endpoint_remove_projectCmd)
}
