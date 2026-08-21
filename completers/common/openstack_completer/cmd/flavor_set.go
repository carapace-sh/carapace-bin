package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flavor_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set flavor properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flavor_setCmd).Standalone()

	flavor_setCmd.Flags().String("description", "", "Set description for the flavor.(Supported")
	flavor_setCmd.Flags().Bool("no-property", false, "Remove all properties from this flavor (specify both --no-property and --property to remove the current properties before setting new properties.)")
	flavor_setCmd.Flags().String("project", "", "Set flavor access to project (name or ID) (admin only)")
	flavor_setCmd.Flags().String("project-domain", "", "Domain the project belongs to (name or ID).")
	flavor_setCmd.Flags().String("property", "", "Property to add or modify for this flavor (repeat option to set multiple properties)")
	flavorCmd.AddCommand(flavor_setCmd)
}
