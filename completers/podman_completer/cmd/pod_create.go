package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pod_createCmd = &cobra.Command{
	Use:   "create [options] [NAME]",
	Short: "Create a new empty pod",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pod_createCmd).Standalone()

	pod_createCmd.Flags().StringSlice("add-host", []string{}, "Add a custom host-to-IP mapping (host:ip) (default [])")
	pod_createCmd.Flags().String("blkio-weight", "", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.")
	pod_createCmd.Flags().StringSlice("blkio-weight-device", []string{}, "Block IO weight (relative device weight, format: `DEVICE_NAME:WEIGHT`)")
	pod_createCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	pod_createCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	pod_createCmd.Flags().String("cpus", "", "Number of CPUs. The default is 0.000 which means no limit")
	pod_createCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	pod_createCmd.Flags().String("cpuset-mems", "", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.")
	pod_createCmd.Flags().StringSlice("device", []string{}, "Add a host device to the container")
	pod_createCmd.Flags().StringSlice("device-read-bps", []string{}, "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)")
	pod_createCmd.Flags().StringSlice("device-write-bps", []string{}, "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)")
	pod_createCmd.Flags().StringSlice("dns", []string{}, "Set custom DNS servers")
	pod_createCmd.Flags().StringSlice("dns-option", []string{}, "Set custom DNS options")
	pod_createCmd.Flags().StringSlice("dns-search", []string{}, "Set custom DNS search domains")
	pod_createCmd.Flags().String("exit-policy", "", "Behaviour when the last container exits")
	pod_createCmd.Flags().StringSlice("gidmap", []string{}, "GID map to use for the user namespace")
	pod_createCmd.Flags().StringSlice("gpus", []string{}, "GPU devices to add to the container ('all' to pass all GPUs)")
	pod_createCmd.Flags().Bool("help", false, "")
	pod_createCmd.Flags().StringP("hostname", "h", "", "Set container hostname")
	pod_createCmd.Flags().String("hosts-file", "", "Base file to create the /etc/hosts file inside the container, or one of the special values. (\"image\"|\"none\")")
	pod_createCmd.Flags().Bool("infra", false, "Create an infra container associated with the pod to share namespaces with")
	pod_createCmd.Flags().String("infra-command", "", "Overwrite the default ENTRYPOINT of the image")
	pod_createCmd.Flags().String("infra-conmon-pidfile", "", "Path to the file that will receive the PID of conmon")
	pod_createCmd.Flags().String("infra-image", "", "Image to use to override builtin infra container")
	pod_createCmd.Flags().String("infra-name", "", "Assign a name to the container")
	pod_createCmd.Flags().String("ip", "", "Specify a static IPv4 address for the container")
	pod_createCmd.Flags().String("ip6", "", "Specify a static IPv6 address for the container")
	pod_createCmd.Flags().StringSliceP("label", "l", []string{}, "Set metadata on container")
	pod_createCmd.Flags().StringSlice("label-file", []string{}, "Read in a line delimited file of labels")
	pod_createCmd.Flags().String("mac-address", "", "Container MAC address (e.g. 92:d0:c6:0a:29:33)")
	pod_createCmd.Flags().StringP("memory", "m", "", "Memory limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	pod_createCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	pod_createCmd.Flags().StringP("name", "n", "", "Assign a name to the pod")
	pod_createCmd.Flags().StringSlice("network", []string{}, "Connect a container to a network")
	pod_createCmd.Flags().StringSlice("network-alias", []string{}, "Add network-scoped alias for the container")
	pod_createCmd.Flags().Bool("no-hostname", false, "Do not create /etc/hostname within the container, instead use the version from the image")
	pod_createCmd.Flags().Bool("no-hosts", false, "Do not create /etc/hosts within the container, instead use the version from the image")
	pod_createCmd.Flags().String("pid", "", "PID namespace to use")
	pod_createCmd.Flags().String("pod-id-file", "", "Write the pod ID to the file")
	pod_createCmd.Flags().StringSliceP("publish", "p", []string{}, "Publish a container's port, or a range of ports, to the host (default [])")
	pod_createCmd.Flags().Bool("replace", false, "If a pod with the same name exists, replace it")
	pod_createCmd.Flags().String("restart", "", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")")
	pod_createCmd.Flags().StringSlice("security-opt", []string{}, "Security Options")
	pod_createCmd.Flags().String("share", "", "A comma delimited list of kernel namespaces the pod will share")
	pod_createCmd.Flags().Bool("share-parent", false, "Set the pod's cgroup as the cgroup parent for all containers joining the pod")
	pod_createCmd.Flags().String("shm-size", "", "Size of /dev/shm (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	pod_createCmd.Flags().String("shm-size-systemd", "", "Size of systemd specific tmpfs mounts (/run, /run/lock) (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	pod_createCmd.Flags().String("subgidname", "", "Name of range listed in /etc/subgid for use in user namespace")
	pod_createCmd.Flags().String("subuidname", "", "Name of range listed in /etc/subuid for use in user namespace")
	pod_createCmd.Flags().StringSlice("sysctl", []string{}, "Sysctl options")
	pod_createCmd.Flags().StringSlice("uidmap", []string{}, "UID map to use for the user namespace")
	pod_createCmd.Flags().String("userns", "", "User namespace to use")
	pod_createCmd.Flags().String("uts", "", "UTS namespace to use")
	pod_createCmd.Flags().StringSliceP("volume", "v", []string{}, "Bind mount a volume into the container")
	pod_createCmd.Flags().StringSlice("volumes-from", []string{}, "Mount volumes from the specified container(s)")
	podCmd.AddCommand(pod_createCmd)
}
