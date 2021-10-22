package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_neighborsCmd = &cobra.Command{
	Use:   "neighbors",
	Short: "List a Droplet's neighbors on your account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_neighborsCmd).Standalone()
	compute_droplet_neighborsCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `PublicIPv4`, `PrivateIPv4`, `PublicIPv6`, `Memory`, `VCPUs`, `Disk`, `Region`, `Image`, `VPCUUID`, `Status`, `Tags`, `Features`, `Volumes`")
	compute_droplet_neighborsCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_dropletCmd.AddCommand(compute_droplet_neighborsCmd)

	carapace.Gen(compute_droplet_neighborsCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "PublicIPv4", "PrivateIPv4", "PublicIPv6", "Memory", "VCPUs", "Disk", "Region", "Image", "VPCUUID", "Status", "Tags", "Features", "Volumes").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
