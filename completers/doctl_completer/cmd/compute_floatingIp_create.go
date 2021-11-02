package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_floatingIp_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new floating IP address",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_floatingIp_createCmd).Standalone()
	compute_floatingIp_createCmd.Flags().Int("droplet-id", 0, "The ID of the Droplet to assign the floating IP to (mutually exclusive with `--region`).")
	compute_floatingIp_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `IP`, `Region`, `DropletID`, `DropletName`")
	compute_floatingIp_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_floatingIp_createCmd.Flags().String("region", "", "Region where to create the floating IP address. (mutually exclusive with `--droplet-id`)")
	compute_floatingIpCmd.AddCommand(compute_floatingIp_createCmd)

	// TODO flag completion
	carapace.Gen(compute_floatingIp_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("IP", "Region", "DropletID", "DropletName").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region": action.ActionRegions(),
	})
}
