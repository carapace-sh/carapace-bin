package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewall_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new cloud firewall",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewall_createCmd).Standalone()
	compute_firewall_createCmd.Flags().StringSlice("droplet-ids", []string{}, "A comma-separated list of Droplet IDs to place behind the cloud firewall, e.g.: `123,456`")
	compute_firewall_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Status`, `Created`, `InboundRules`, `OutboundRules`, `DropletIDs`, `Tags`, `PendingChanges`")
	compute_firewall_createCmd.Flags().String("inbound-rules", "", "A comma-separated key-value list that defines an inbound rule, e.g.: `protocol:tcp,ports:22,droplet_id:123`. Use a quoted string of space-separated values for multiple rules.")
	compute_firewall_createCmd.Flags().String("name", "", "Firewall name (required)")
	compute_firewall_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_firewall_createCmd.Flags().String("outbound-rules", "", "A comma-separate key-value list the defines an outbound rule, e.g.: `protocol:tcp,ports:22,address:0.0.0.0/0`. Use a quoted string of space-separated values for multiple rules.")
	compute_firewall_createCmd.Flags().StringSlice("tag-names", []string{}, "A comma-separated list of tag names to apply to the cloud firewall, e.g.: `frontend,backend`")
	compute_firewallCmd.AddCommand(compute_firewall_createCmd)

	// TODO flag completion
	carapace.Gen(compute_firewall_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Status", "Created", "InboundRules", "OutboundRules", "DropletIDs", "Tags", "PendingChanges").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
