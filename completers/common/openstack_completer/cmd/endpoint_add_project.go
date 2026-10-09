package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_add_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Associate a project to an endpoint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_add_projectCmd).Standalone()

	endpoint_add_projectCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	endpoint_addCmd.AddCommand(endpoint_add_projectCmd)
}
