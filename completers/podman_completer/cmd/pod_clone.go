package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_cloneCmd = &cobra.Command{
	Use:   "clone [options] POD NAME",
	Short: "Clone an existing pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_cloneCmd).Standalone()

	pod_cloneCmd.Flags().String("blkio-weight", "", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.")
	pod_cloneCmd.Flags().StringSlice("blkio-weight-device", []string{}, "Block IO weight (relative device weight, format: `DEVICE_NAME:WEIGHT`)")
	pod_cloneCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	pod_cloneCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	pod_cloneCmd.Flags().String("cpus", "", "Number of CPUs. The default is 0.000 which means no limit")
	pod_cloneCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	pod_cloneCmd.Flags().String("cpuset-mems", "", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.")
	pod_cloneCmd.Flags().Bool("destroy", false, "destroy the original pod")
	pod_cloneCmd.Flags().StringSlice("device", []string{}, "Add a host device to the container")
	pod_cloneCmd.Flags().StringSlice("device-read-bps", []string{}, "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)")
	pod_cloneCmd.Flags().StringSlice("device-write-bps", []string{}, "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)")
	pod_cloneCmd.Flags().StringSlice("gidmap", []string{}, "GID map to use for the user namespace")
	pod_cloneCmd.Flags().StringSlice("gpus", []string{}, "GPU devices to add to the container ('all' to pass all GPUs)")
	pod_cloneCmd.Flags().Bool("help", false, "")
	pod_cloneCmd.Flags().StringP("hostname", "h", "", "Set container hostname")
	pod_cloneCmd.Flags().String("infra-command", "", "Overwrite the default ENTRYPOINT of the image")
	pod_cloneCmd.Flags().String("infra-conmon-pidfile", "", "Path to the file that will receive the PID of conmon")
	pod_cloneCmd.Flags().String("infra-name", "", "Assign a name to the container")
	pod_cloneCmd.Flags().StringSliceP("label", "l", []string{}, "Set metadata on container")
	pod_cloneCmd.Flags().StringSlice("label-file", []string{}, "Read in a line delimited file of labels")
	pod_cloneCmd.Flags().StringP("memory", "m", "", "Memory limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	pod_cloneCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	pod_cloneCmd.Flags().StringP("name", "n", "", "name the new pod")
	pod_cloneCmd.Flags().String("pid", "", "PID namespace to use")
	pod_cloneCmd.Flags().String("restart", "", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")")
	pod_cloneCmd.Flags().StringSlice("security-opt", []string{}, "Security Options")
	pod_cloneCmd.Flags().String("shm-size", "", "Size of /dev/shm (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	pod_cloneCmd.Flags().String("shm-size-systemd", "", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	pod_cloneCmd.Flags().Bool("start", false, "start the new pod")
	pod_cloneCmd.Flags().String("subgidname", "", "Name of range listed in /etc/subgid for use in user namespace")
	pod_cloneCmd.Flags().String("subuidname", "", "Name of range listed in /etc/subuid for use in user namespace")
	pod_cloneCmd.Flags().StringSlice("sysctl", []string{}, "Sysctl options")
	pod_cloneCmd.Flags().StringSlice("uidmap", []string{}, "UID map to use for the user namespace")
	pod_cloneCmd.Flags().String("userns", "", "User namespace to use")
	pod_cloneCmd.Flags().String("uts", "", "UTS namespace to use")
	pod_cloneCmd.Flags().StringSliceP("volume", "v", []string{}, "Bind mount a volume into the container")
	pod_cloneCmd.Flags().StringSlice("volumes-from", []string{}, "Mount volumes from the specified container(s)")
	podCmd.AddCommand(pod_cloneCmd)
}
