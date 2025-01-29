package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var machine_setCmd = &cobra.Command{
	Use:   "set [options] [NAME]",
	Short: "Set a virtual machine setting",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(machine_setCmd).Standalone()

	machine_setCmd.Flags().String("cpus", "", "Number of CPUs")
	machine_setCmd.Flags().String("disk-size", "", "Disk size in GiB")
	machine_setCmd.Flags().StringP("memory", "m", "", "Memory in MiB")
	machine_setCmd.Flags().Bool("rootful", false, "Whether this machine should prefer rootful container execution")
	machine_setCmd.Flags().StringSlice("usb", []string{}, "USBs bus=$1,devnum=$2 or vendor=$1,product=$2")
	machine_setCmd.Flags().Bool("user-mode-networking", false, "Whether this machine should use user-mode networking, routing traffic through a host user-space process")
	machineCmd.AddCommand(machine_setCmd)
}
