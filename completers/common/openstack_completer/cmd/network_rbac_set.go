package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/openstack_completer/cmd/action"
	"github.com/spf13/cobra"
)

var network_rbac_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set network RBAC policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_rbac_setCmd).Standalone()

	network_rbac_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_rbac_setCmd.Flags().String("target-project", "", "The project to which the RBAC policy will be enforced (name or ID)")
	network_rbac_setCmd.Flags().String("target-project-domain", "", "Domain the target project belongs to (name or ID).")
	network_rbacCmd.AddCommand(network_rbac_setCmd)

	carapace.Gen(network_rbac_setCmd).FlagCompletion(carapace.ActionMap{
		"target-project":        action.ActionProjects(),
		"target-project-domain": action.ActionDomains(),
	})

	carapace.Gen(network_rbac_setCmd).PositionalCompletion(action.ActionNetworkRbacPolicies())
}
