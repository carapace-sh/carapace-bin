package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var floating_ip_unsetCmd = &cobra.Command{
	Use:   "unset",
	Short: "Unset floating IP Properties",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(floating_ip_unsetCmd).Standalone()

	floating_ip_unsetCmd.Flags().Bool("all-tag", false, "Clear all tags associated with the floating IP")
	floating_ip_unsetCmd.Flags().String("extra-property", "", "Additional parameters can be passed using this property.")
	floating_ip_unsetCmd.Flags().Bool("port", false, "Disassociate any port associated with the floating IP")
	floating_ip_unsetCmd.Flags().Bool("qos-policy", false, "Remove the QoS policy attached to the floating IP")
	floating_ip_unsetCmd.Flags().String("tag", "", "Tag to be removed from the floating IP (repeat option to remove multiple tags)")
	floating_ipCmd.AddCommand(floating_ip_unsetCmd)
}
