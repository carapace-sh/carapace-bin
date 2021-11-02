package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/doctl_completer/cmd/action"
	"github.com/spf13/cobra"
)

var compute_droplet_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Droplet",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(compute_droplet_createCmd).Standalone()
	compute_droplet_createCmd.Flags().Bool("droplet-agent", false, "By default, the agent is installed on new Droplets but installation errors are ignored. Set --droplet-agent=false to prevent installation. Set `true` to make installation errors fatal.")
	compute_droplet_createCmd.Flags().Bool("enable-backups", false, "Enables backups for the Droplet")
	compute_droplet_createCmd.Flags().Bool("enable-ipv6", false, "Enables IPv6 support and assigns an IPv6 address")
	compute_droplet_createCmd.Flags().Bool("enable-monitoring", false, "Install the DigitalOcean agent for additional monitoring")
	compute_droplet_createCmd.Flags().Bool("enable-private-networking", false, "Enables private networking for the Droplet by provisioning it inside of your account's default VPC for the region")
	compute_droplet_createCmd.Flags().String("format", "", "Columns for output in a comma-separated list. Possible values: `ID`, `Name`, `PublicIPv4`, `PrivateIPv4`, `PublicIPv6`, `Memory`, `VCPUs`, `Disk`, `Region`, `Image`, `VPCUUID`, `Status`, `Tags`, `Features`, `Volumes`")
	compute_droplet_createCmd.Flags().String("image", "", "An ID or slug indicating the image the Droplet will be based-on (e.g. `ubuntu-20-04-x64`). Use the commands under `doctl compute image` to find additional images. (required)")
	compute_droplet_createCmd.Flags().Bool("no-header", false, "Return raw data with no headers")
	compute_droplet_createCmd.Flags().String("region", "", "A slug indicating the region where the Droplet will be created (e.g. `nyc1`). Run `doctl compute region list` for a list of valid regions. (required)")
	compute_droplet_createCmd.Flags().String("size", "", "A slug indicating the size of the Droplet (e.g. `s-1vcpu-1gb`). Run `doctl compute size list` for a list of valid sizes. (required)")
	compute_droplet_createCmd.Flags().StringSlice("ssh-keys", []string{}, "A list of SSH key fingerprints or IDs of the SSH keys to embed in the Droplet's root account upon creation")
	compute_droplet_createCmd.Flags().String("tag-name", "", "A tag name to be applied to the Droplet")
	compute_droplet_createCmd.Flags().StringSlice("tag-names", []string{}, "A list of tag names to be applied to the Droplet")
	compute_droplet_createCmd.Flags().String("user-data", "", "User-data to configure the Droplet on first boot")
	compute_droplet_createCmd.Flags().String("user-data-file", "", "The path to a file containing user-data to configure the Droplet on first boot")
	compute_droplet_createCmd.Flags().StringSlice("volumes", []string{}, "A list of block storage volume IDs to attach to the Droplet")
	compute_droplet_createCmd.Flags().String("vpc-uuid", "", "The UUID of a non-default VPC to create the Droplet in")
	compute_droplet_createCmd.Flags().Bool("wait", false, "Wait for Droplet creation to complete before returning")
	compute_dropletCmd.AddCommand(compute_droplet_createCmd)

	// TODO flag completion
	carapace.Gen(compute_droplet_createCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.ActionValues("ID", "Name", "PublicIPv4", "PrivateIPv4", "PublicIPv6", "Memory", "VCPUs", "Disk", "Region", "Image", "VPCUUID", "Status", "Tags", "Features", "Volumes").Invoke(c).Filter(c.Parts).ToA()
		}),
		"region":         action.ActionRegions(),
		"user-data-file": carapace.ActionFiles(),
	})
}
