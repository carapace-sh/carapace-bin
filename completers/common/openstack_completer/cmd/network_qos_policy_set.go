package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var network_qos_policy_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set QoS policy properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(network_qos_policy_setCmd).Standalone()

	network_qos_policy_setCmd.Flags().Bool("default", false, "Set this as a default network QoS policy")
	network_qos_policy_setCmd.Flags().String("description", "", "Description of the QoS policy")
	network_qos_policy_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	network_qos_policy_setCmd.Flags().String("name", "", "Set QoS policy name")
	network_qos_policy_setCmd.Flags().Bool("no-default", false, "Set this as a non-default network QoS policy")
	network_qos_policy_setCmd.Flags().Bool("no-share", false, "Make the QoS policy not accessible by other projects")
	network_qos_policy_setCmd.Flags().Bool("share", false, "Make the QoS policy accessible by other projects")
	network_qos_policyCmd.AddCommand(network_qos_policy_setCmd)
}
