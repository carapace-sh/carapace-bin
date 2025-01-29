package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_initCmd = &cobra.Command{
	Use:   "init [options] [NAME]",
	Short: "Initialize a virtual machine",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_initCmd).Standalone()

	machine_initCmd.Flags().String("cpus", "", "Number of CPUs")
	machine_initCmd.Flags().String("disk-size", "", "Disk size in GiB")
	machine_initCmd.Flags().String("ignition-path", "", "Path to ignition file")
	machine_initCmd.Flags().String("image", "", "Bootable image for machine")
	machine_initCmd.Flags().String("image-path", "", "Bootable image for machine")
	machine_initCmd.Flags().StringP("memory", "m", "", "Memory in MiB")
	machine_initCmd.Flags().Bool("now", false, "Start machine now")
	machine_initCmd.Flags().Bool("reexec", false, "process was rexeced")
	machine_initCmd.Flags().Bool("rootful", false, "Whether this machine should prefer rootful container execution")
	machine_initCmd.Flags().String("timezone", "", "Set timezone")
	machine_initCmd.Flags().StringSlice("usb", []string{}, "USB Host passthrough: bus=$1,devnum=$2 or vendor=$1,product=$2")
	machine_initCmd.Flags().Bool("user-mode-networking", false, "Whether this machine should use user-mode networking, routing traffic through a host user-space process")
	machine_initCmd.Flags().String("username", "", "Username used in image")
	machine_initCmd.Flags().StringSliceP("volume", "v", []string{}, "Volumes to mount, source:target")
	machine_initCmd.Flags().String("volume-driver", "", "Optional volume driver")
	machine_initCmd.Flag("image-path").Hidden = true
	machine_initCmd.Flag("reexec").Hidden = true
	machine_initCmd.Flag("volume-driver").Hidden = true
	machineCmd.AddCommand(machine_initCmd)
}
