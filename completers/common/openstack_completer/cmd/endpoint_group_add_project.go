package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_group_add_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Add a project to an endpoint group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_group_add_projectCmd).Standalone()

	endpoint_group_add_projectCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	endpoint_group_addCmd.AddCommand(endpoint_group_add_projectCmd)
}
