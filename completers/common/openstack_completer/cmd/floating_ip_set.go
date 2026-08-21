package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set floating IP Properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_setCmd).Standalone()

	floating_ip_setCmd.Flags().String("description", "", "Set floating IP description")
	floating_ip_setCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	floating_ip_setCmd.Flags().String("fixed-ip-address", "", "Fixed IP of the port (required only if port has multiple IPs)")
	floating_ip_setCmd.Flags().Bool("no-qos-policy", false, "Remove the QoS policy attached to the floating IP")
	floating_ip_setCmd.Flags().Bool("no-tag", false, "Clear tags associated with the floating IP.")
	floating_ip_setCmd.Flags().String("port", "", "Associate the floating IP with port (name or ID)")
	floating_ip_setCmd.Flags().String("qos-policy", "", "Attach QoS policy to the floating IP (name or ID)")
	floating_ip_setCmd.Flags().String("tag", "", "Tag to be added to the floating IP (repeat option to set multiple tags)")
	floating_ipCmd.AddCommand(floating_ip_setCmd)
}
