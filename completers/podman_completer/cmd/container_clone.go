package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_cloneCmd = &cobra.Command{
	Use:   "clone [options] CONTAINER NAME IMAGE",
	Short: "Clone an existing container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_cloneCmd).Standalone()

	container_cloneCmd.Flags().String("blkio-weight", "", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.")
	container_cloneCmd.Flags().StringSlice("blkio-weight-device", []string{}, "Block IO weight (relative device weight, format: `DEVICE_NAME:WEIGHT`)")
	container_cloneCmd.Flags().String("cpu-period", "", "Limit the CPU CFS (Completely Fair Scheduler) period")
	container_cloneCmd.Flags().String("cpu-quota", "", "Limit the CPU CFS (Completely Fair Scheduler) quota")
	container_cloneCmd.Flags().String("cpu-rt-period", "", "Limit the CPU real-time period in microseconds")
	container_cloneCmd.Flags().String("cpu-rt-runtime", "", "Limit the CPU real-time runtime in microseconds")
	container_cloneCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	container_cloneCmd.Flags().String("cpus", "", "Number of CPUs. The default is 0.000 which means no limit")
	container_cloneCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	container_cloneCmd.Flags().String("cpuset-mems", "", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.")
	container_cloneCmd.Flags().Bool("destroy", false, "destroy the original container")
	container_cloneCmd.Flags().StringSlice("device-read-bps", []string{}, "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)")
	container_cloneCmd.Flags().StringSlice("device-write-bps", []string{}, "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)")
	container_cloneCmd.Flags().BoolP("force", "f", false, "force the existing container to be destroyed")
	container_cloneCmd.Flags().StringP("memory", "m", "", "Memory limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	container_cloneCmd.Flags().String("memory-reservation", "", "Memory soft limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	container_cloneCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	container_cloneCmd.Flags().String("memory-swappiness", "", "Tune container memory swappiness (0 to 100, or -1 for system default)")
	container_cloneCmd.Flags().String("name", "", "Assign a name to the container")
	container_cloneCmd.Flags().String("pod", "", "Run container in an existing pod")
	container_cloneCmd.Flags().Bool("run", false, "run the new container")
	containerCmd.AddCommand(container_cloneCmd)
}
