package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_droplet_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Droplets on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_listCmd).Standalone()
	compute_droplet_listCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `PublicIPv4`, `PrivateIPv4`, `PublicIPv6`, `Memory`, `VCPUs`, `Disk`, `Region`, `Image`, `VPCUUID`, `Status`, `Tags`, `Features`, `Volumes`")
	compute_droplet_listCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_droplet_listCmd.Flags().String("region", "", "Droplet region")
	compute_droplet_listCmd.Flags().String("tag-name", "", "Tag name")
	compute_dropletCmd.AddCommand(compute_droplet_listCmd)

	carapace.Gen(compute_droplet_listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "PublicIPv4", "PrivateIPv4", "PublicIPv6", "Memory", "VCPUs", "Disk", "Region", "Image", "VPCUUID", "Status", "Tags", "Features", "Volumes").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region": action.ActionRegions(),
	})
}
