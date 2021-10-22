package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var compute_droplet_getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve information about a Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_getCmd).Standalone()
	compute_droplet_getCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `PublicIPv4`, `PrivateIPv4`, `PublicIPv6`, `Memory`, `VCPUs`, `Disk`, `Region`, `Image`, `VPCUUID`, `Status`, `Tags`, `Features`, `Volumes`")
	compute_droplet_getCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_droplet_getCmd.Flags().String("template", "", "Go template format. Sample values: `{{.ID}}`, `{{.Name}}`, `{{.Memory}}`, `{{.Region.Name}}`, `{{.Image}}`, `{{.Tags}}`")
	compute_dropletCmd.AddCommand(compute_droplet_getCmd)

	carapace.Gen(compute_droplet_getCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "PublicIPv4", "PrivateIPv4", "PublicIPv6", "Memory", "VCPUs", "Disk", "Region", "Image", "VPCUUID", "Status", "Tags", "Features", "Volumes").Invoke(c).Filter(c.Parts).ToA()
		}),
	})
}
