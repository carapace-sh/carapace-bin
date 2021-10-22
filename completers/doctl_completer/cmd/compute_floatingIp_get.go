package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_floatingIp_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a floating IP address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIp_getCmd).Standalone()
	compute_floatingIp_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `IP`, `Region`, `DropletID`, `DropletName`")
	compute_floatingIp_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_floatingIpCmd.AddCommand(compute_floatingIp_getCmd)

	carapace.Gen(compute_floatingIp_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("IP", "Region", "DropletID", "DropletName").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
