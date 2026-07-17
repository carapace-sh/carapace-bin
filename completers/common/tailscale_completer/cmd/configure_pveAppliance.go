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

	configure_pveApplianceCmd.Flags().String("bridge", "", "bridge to attach the VM to")
	configure_pveApplianceCmd.Flags().Int("cores", 0, "number of CPU cores")
	configure_pveApplianceCmd.Flags().String("disk-size", "", "disk size for the VM")
	configure_pveApplianceCmd.Flags().Int("memory", 0, "memory in MB for the VM")
	configure_pveApplianceCmd.Flags().String("storage", "", "storage name for the VM disk")
	configure_pveApplianceCmd.Flags().String("variant", "", "appliance variant (pi-arm64, vm-amd64, vm-arm64)")
	configure_pveApplianceCmd.Flags().Int("vmid", 0, "VM ID for the new Proxmox VE VM")
	configureCmd.AddCommand(configure_pveApplianceCmd)
}
