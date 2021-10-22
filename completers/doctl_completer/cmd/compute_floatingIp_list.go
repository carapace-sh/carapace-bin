package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_floatingIp_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all floating IP addresses on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIp_listCmd).Standalone()
	compute_floatingIp_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `IP`, `Region`, `DropletID`, `DropletName`")
	compute_floatingIp_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_floatingIp_listCmd.Flags().String("region", "", "The region the floating IP address resides in")
	compute_floatingIpCmd.AddCommand(compute_floatingIp_listCmd)

	carapace.Gen(compute_floatingIp_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("IP", "Region", "DropletID", "DropletName").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region": action.ActionRegions(),
	})
}
