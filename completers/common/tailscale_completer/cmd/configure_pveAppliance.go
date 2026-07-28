package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_pveApplianceCmd = &cobra.Command{
	Use:   "pve-appliance",
	Short: "Create a Proxmox VE VM running the Tailscale appliance image [experimental]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_pveApplianceCmd).Standalone()

	configure_pveApplianceCmd.Flags().String("add-ssh-authorized-keys", "", "path to an authorized_keys file to include on the appliance for breakglass SSH access")
	configure_pveApplianceCmd.Flags().String("bridge", "", "bridge to attach the VM to")
	configure_pveApplianceCmd.Flags().Int("cores", 2, "number of CPU cores")
	configure_pveApplianceCmd.Flags().String("disk-size", "4G", "disk size for the VM (accepts K/M/G suffixes)")
	configure_pveApplianceCmd.Flags().String("gaf", "", "use a local GAF file instead of downloading (skips signature verification)")
	configure_pveApplianceCmd.Flags().Int("memory", 1024, "memory in MB for the VM")
	configure_pveApplianceCmd.Flags().String("name", "", "VM name; defaults to tsapp-<vmid>")
	configure_pveApplianceCmd.Flags().Bool("start", true, "start the VM after import")
	configure_pveApplianceCmd.Flags().String("storage", "", "PVE storage to import the disk into (e.g. local-lvm)")
	configure_pveApplianceCmd.Flags().String("track", "", "which track to download from; defaults to the current track")
	configure_pveApplianceCmd.Flags().String("variant", "vm-amd64", "appliance variant (vm-amd64, vm-arm64)")
	configure_pveApplianceCmd.Flags().Int("vmid", 0, "VM ID for the new Proxmox VE VM")
	configure_pveApplianceCmd.Flags().Bool("yes", false, "skip the confirmation prompt")
	configureCmd.AddCommand(configure_pveApplianceCmd)

	carapace.Gen(configure_pveApplianceCmd).FlagCompletion(carapace.ActionMap{
		"add-ssh-authorized-keys": carapace.ActionFiles(),
		"gaf":                     carapace.ActionFiles(),
		"track":                   carapace.ActionValues("stable", "release-candidate", "unstable"),
		"variant":                 carapace.ActionValues("vm-amd64", "vm-arm64"),
	})
}
