package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:     "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short:   "Create and run a new container from an image",
	GroupID: "common",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().String("add-host", "", "Add a custom host-to-IP mapping (host:ip)")
	runCmd.Flags().StringP("attach", "a", "", "Attach to STDIN, STDOUT or STDERR")
	runCmd.Flags().Uint16("blkio-weight", 0, "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)")
	runCmd.Flags().String("blkio-weight-device", "", "Block IO weight (relative device weight)")
	runCmd.Flags().String("cap-add", "", "Add Linux capabilities")
	runCmd.Flags().String("cap-drop", "", "Drop Linux capabilities")
	runCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	runCmd.Flags().String("cgroupns", "", "Cgroup namespace to use (host|private)")
	runCmd.Flags().String("cidfile", "", "Write the container ID to the file")
	runCmd.Flags().Int64("cpu-count", 0, "CPU count (Windows only)")
	runCmd.Flags().Int64("cpu-percent", 0, "CPU percent (Windows only)")
	runCmd.Flags().Int64("cpu-period", 0, "Limit CPU CFS (Completely Fair Scheduler) period")
	runCmd.Flags().Int64("cpu-quota", 0, "Limit CPU CFS (Completely Fair Scheduler) quota")
	runCmd.Flags().Int64("cpu-rt-period", 0, "Limit CPU real-time period in microseconds")
	runCmd.Flags().Int64("cpu-rt-runtime", 0, "Limit CPU real-time runtime in microseconds")
	runCmd.Flags().Int64P("cpu-shares", "c", 0, "CPU shares (relative weight)")
	runCmd.Flags().String("cpus", "", "Number of CPUs")
	runCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	runCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	runCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
	runCmd.Flags().String("device", "", "Add a host device to the container")
	runCmd.Flags().String("device-cgroup-rule", "", "Add a rule to the cgroup allowed devices list")
	runCmd.Flags().String("device-read-bps", "", "Limit read rate (bytes per second) from a device")
	runCmd.Flags().String("device-read-iops", "", "Limit read rate (IO per second) from a device")
	runCmd.Flags().String("device-write-bps", "", "Limit write rate (bytes per second) to a device")
	runCmd.Flags().String("device-write-iops", "", "Limit write rate (IO per second) to a device")
	runCmd.Flags().Bool("disable-content-trust", true, "Skip image verification")
	runCmd.Flags().String("dns", "", "Set custom DNS servers")
	runCmd.Flags().String("dns-opt", "", "Set DNS options")
	runCmd.Flags().String("dns-option", "", "Set DNS options")
	runCmd.Flags().String("dns-search", "", "Set custom DNS search domains")
	runCmd.Flags().String("domainname", "", "Container NIS domain name")
	runCmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image")
	runCmd.Flags().StringP("env", "e", "", "Set environment variables")
	runCmd.Flags().String("env-file", "", "Read in a file of environment variables")
	runCmd.Flags().String("expose", "", "Expose a port or a range of ports")
	runCmd.Flags().String("gpus", "", "GPU devices to add to the container ('all' to pass all GPUs)")
	runCmd.Flags().String("group-add", "", "Add additional groups to join")
	runCmd.Flags().String("health-cmd", "", "Command to run to check health")
	runCmd.Flags().Duration("health-interval", 0, "Time between running the check (ms|s|m|h) (default 0s)")
	runCmd.Flags().Int("health-retries", 0, "Consecutive failures needed to report unhealthy")
	runCmd.Flags().Duration("health-start-period", 0, "Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)")
	runCmd.Flags().Duration("health-timeout", 0, "Maximum time to allow one check to run (ms|s|m|h) (default 0s)")
	runCmd.Flags().Bool("help", false, "Print usage")
	runCmd.Flags().StringP("hostname", "h", "", "Container host name")
	runCmd.Flags().Bool("init", false, "Run an init inside the container that forwards signals and reaps processes")
	runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	runCmd.Flags().String("io-maxbandwidth", "", "Maximum IO bandwidth limit for the system drive (Windows only)")
	runCmd.Flags().Uint64("io-maxiops", 0, "Maximum IOps limit for the system drive (Windows only)")
	runCmd.Flags().String("ip", "", "IPv4 address (e.g., 172.30.100.104)")
	runCmd.Flags().String("ip6", "", "IPv6 address (e.g., 2001:db8::33)")
	runCmd.Flags().String("ipc", "", "IPC mode to use")
	runCmd.Flags().String("isolation", "", "Container isolation technology")
	runCmd.Flags().String("kernel-memory", "", "Kernel memory limit")
	runCmd.Flags().StringP("label", "l", "", "Set meta data on a container")
	runCmd.Flags().String("label-file", "", "Read in a line delimited file of labels")
	runCmd.Flags().String("link", "", "Add link to another container")
	runCmd.Flags().String("link-local-ip", "", "Container IPv4/IPv6 link-local addresses")
	runCmd.Flags().String("log-driver", "", "Logging driver for the container")
	runCmd.Flags().String("log-opt", "", "Log driver options")
	runCmd.Flags().String("mac-address", "", "Container MAC address (e.g., 92:d0:c6:0a:29:33)")
	runCmd.Flags().StringP("memory", "m", "", "Memory limit")
	runCmd.Flags().String("memory-reservation", "", "Memory soft limit")
	runCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	runCmd.Flags().Int64("memory-swappiness", 0, "Tune container memory swappiness (0 to 100)")
	runCmd.Flags().String("mount", "", "Attach a filesystem mount to the container")
	runCmd.Flags().String("name", "", "Assign a name to the container")
	runCmd.Flags().String("net", "", "Connect a container to a network")
	runCmd.Flags().String("net-alias", "", "Add network-scoped alias for the container")
	runCmd.Flags().String("network", "", "Connect a container to a network")
	runCmd.Flags().String("network-alias", "", "Add network-scoped alias for the container")
	runCmd.Flags().Bool("no-healthcheck", false, "Disable any container-specified HEALTHCHECK")
	runCmd.Flags().Bool("oom-kill-disable", false, "Disable OOM Killer")
	runCmd.Flags().Int("oom-score-adj", 0, "Tune host's OOM preferences (-1000 to 1000)")
	runCmd.Flags().String("pid", "", "PID namespace to use")
	runCmd.Flags().Int64("pids-limit", 0, "Tune container pids limit (set -1 for unlimited)")
	runCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	runCmd.Flags().Bool("privileged", false, "Give extended privileges to this container")
	runCmd.Flags().StringP("publish", "p", "", "Publish a container's port(s) to the host")
	runCmd.Flags().BoolP("publish-all", "P", false, "Publish all exposed ports to random ports")
	runCmd.Flags().String("pull", "missing", "Pull image before running (\"always\", \"missing\", \"never\")")
	runCmd.Flags().BoolP("quiet", "q", false, "Suppress the pull output")
	runCmd.Flags().Bool("read-only", false, "Mount the container's root filesystem as read only")
	runCmd.Flags().String("restart", "no", "Restart policy to apply when a container exits")
	runCmd.Flags().Bool("rm", false, "Automatically remove the container when it exits")
	runCmd.Flags().String("runtime", "", "Runtime to use for this container")
	runCmd.Flags().String("security-opt", "", "Security Options")
	runCmd.Flags().String("shm-size", "", "Size of /dev/shm")
	runCmd.Flags().Bool("sig-proxy", true, "Proxy received signals to the process")
	runCmd.Flags().String("stop-signal", "", "Signal to stop the container")
	runCmd.Flags().Int("stop-timeout", 0, "Timeout (in seconds) to stop a container")
	runCmd.Flags().String("storage-opt", "", "Storage driver options for the container")
	runCmd.Flags().String("sysctl", "", "Sysctl options")
	runCmd.Flags().String("tmpfs", "", "Mount a tmpfs directory")
	runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	runCmd.Flags().String("ulimit", "", "Ulimit options")
	runCmd.Flags().StringP("user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
	runCmd.Flags().String("userns", "", "User namespace to use")
	runCmd.Flags().String("uts", "", "UTS namespace to use")
	runCmd.Flags().StringP("volume", "v", "", "Bind mount a volume")
	runCmd.Flags().String("volume-driver", "", "Optional volume driver for the container")
	runCmd.Flags().String("volumes-from", "", "Mount volumes from the specified container(s)")
	runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	rootCmd.AddCommand(runCmd)

	// TODO flag completion
	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"attach":      carapace.ActionValues("STDIN", "STDOUT", "STDERR"),
		"cgroupns":    carapace.ActionValues("host", "private"),
		"cidfile":     carapace.ActionFiles(),
		"detach-keys": docker.ActionDetachKeys(),
		"env-file":    carapace.ActionFiles(),
		"label-file":  carapace.ActionFiles(),
		"log-driver":  docker.ActionLogDrivers(),
		"volume":      docker.ActionVolumes(),
		"volumes-from": carapace.Batch(
			docker.ActionContainers(),
			docker.ActionContainerIds(),
		).ToA(),
	})
}
