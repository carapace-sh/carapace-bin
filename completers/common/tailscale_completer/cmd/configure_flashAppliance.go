package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_flashApplianceCmd = &cobra.Command{
	Use:   "flash-appliance",
	Short: "Download a signed Tailscale appliance image and write it to a local disk [experimental]",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_flashApplianceCmd).Standalone()

	configure_flashApplianceCmd.Flags().String("add-ssh-authorized-keys", "", "path to an authorized_keys file to include on the appliance for breakglass SSH access")
	configure_flashApplianceCmd.Flags().String("disk", "", "target block device (e.g. /dev/sdb or /dev/disk4)")
	configure_flashApplianceCmd.Flags().String("gaf", "", "use a local GAF file instead of downloading (skips signature verification)")
	configure_flashApplianceCmd.Flags().String("track", "", "which track to download from; defaults to the current track")
	configure_flashApplianceCmd.Flags().String("variant", "", "appliance variant (pi-arm64, vm-amd64, vm-arm64)")
	configure_flashApplianceCmd.Flags().Bool("yes", false, "skip the destructive-write confirmation prompt")
	configureCmd.AddCommand(configure_flashApplianceCmd)

	carapace.Gen(configure_flashApplianceCmd).FlagCompletion(carapace.ActionMap{
		"add-ssh-authorized-keys": carapace.ActionFiles(),
		"disk":                    carapace.ActionFiles(),
		"gaf":                     carapace.ActionFiles(),
		"track":                   carapace.ActionValues("stable", "release-candidate", "unstable"),
		"variant":                 carapace.ActionValues("pi-arm64", "vm-amd64", "vm-arm64"),
	})
}
