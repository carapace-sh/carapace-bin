package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
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

	carapace.Gen(flavor_unsetCmd).FlagCompletion(carapace.ActionMap{
		"project":        action.ActionProjects(flavor_unsetCmd),
		"project-domain": action.ActionDomains(flavor_unsetCmd),
	})

	carapace.Gen(flavor_unsetCmd).PositionalCompletion(action.ActionFlavors(flavor_unsetCmd))
}
