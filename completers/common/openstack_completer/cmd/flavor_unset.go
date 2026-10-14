package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flavor_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset flavor properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavor_unsetCmd).Standalone()

	flavor_unsetCmd.Flags().String("project", "", "Remove flavor access from project (name or ID) (admin only)")
	flavor_unsetCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	flavor_unsetCmd.Flags().String("property", "", "Property to remove from flavor (repeat option to unset multiple properties)")
	flavorCmd.AddCommand(flavor_unsetCmd)
}
