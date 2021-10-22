package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_firewall_listByDropletCmd = &cobra.Command{
	Use:   "list-by-droplet",
	Short: "List firewalls by Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_firewall_listByDropletCmd).Standalone()
	compute_firewall_listByDropletCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `Status`, `Created`, `InboundRules`, `OutboundRules`, `DropletIDs`, `Tags`, `PendingChanges`")
	compute_firewall_listByDropletCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_firewallCmd.AddCommand(compute_firewall_listByDropletCmd)

	carapace.Gen(compute_firewall_listByDropletCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "Status", "Created", "InboundRules", "OutboundRules", "DropletIDs", "Tags", "PendingChanges").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
