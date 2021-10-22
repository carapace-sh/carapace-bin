package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewall_addRulesCmd = &cobra.Command{
	Use:   "add-rules",
	Short: "Add inbound or outbound rules to a cloud firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewall_addRulesCmd).Standalone()
	compute_firewall_addRulesCmd.Flags().String("inbound-rules", "", "A comma-separated key-value list that defines an inbound rule, e.g.: `protocol:tcp,ports:22,droplet_id:123`. Use a quoted string of space-separated values for multiple rules.")
	compute_firewall_addRulesCmd.Flags().String("outbound-rules", "", "A comma-separate key-value list the defines an outbound rule, e.g.: `protocol:tcp,ports:22,address:0.0.0.0/0`. Use a quoted string of space-separated values for multiple rules.")
	compute_firewallCmd.AddCommand(compute_firewall_addRulesCmd)
}
