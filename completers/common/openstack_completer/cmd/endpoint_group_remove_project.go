package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var endpoint_group_remove_projectCmd = &cobra.Command{
	Use:   "project",
	Short: "Remove project from endpoint group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(endpoint_group_remove_projectCmd).Standalone()

	endpoint_group_remove_projectCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	endpoint_group_removeCmd.AddCommand(endpoint_group_remove_projectCmd)
}
