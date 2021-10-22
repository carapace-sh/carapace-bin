package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_loadBalancer_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve a load balancer",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_loadBalancer_getCmd).Standalone()
	compute_loadBalancer_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `IP`, `Name`, `Status`, `Created`, `Algorithm`, `Region`, `Size`, `SizeUnit`, `VPCUUID`, `Tag`, `DropletIDs`, `RedirectHttpToHttps`, `StickySessions`, `HealthCheck`, `ForwardingRules`, `DisableLetsEncryptDNSRecords`")
	compute_loadBalancer_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_loadBalancerCmd.AddCommand(compute_loadBalancer_getCmd)

	carapace.Gen(compute_loadBalancer_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "IP", "Name", "Status", "Created", "Algorithm", "Region", "Size", "SizeUnit", "VPCUUID", "Tag", "DropletIDs", "RedirectHttpToHttps", "StickySessions", "HealthCheck", "ForwardingRules", "DisableLetsEncryptDNSRecords").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
