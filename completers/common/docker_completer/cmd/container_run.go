package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_runCmd = &cobra.Command{
	Use:   "run [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: "Create and run a new container from an image",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_runCmd).Standalone()

	container_runCmd.Flags().String("add-host", "", "Add a custom host-to-IP mapping (host:ip)")
	container_runCmd.Flags().StringP("attach", "a", "", "Attach to STDIN, STDOUT or STDERR")
	container_runCmd.Flags().Uint16("blkio-weight", 0, "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)")
	container_runCmd.Flags().String("blkio-weight-device", "", "Block IO weight (relative device weight)")
	container_runCmd.Flags().String("cap-add", "", "Add Linux capabilities")
	container_runCmd.Flags().String("cap-drop", "", "Drop Linux capabilities")
	container_runCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	container_runCmd.Flags().String("cgroupns", "", "Cgroup namespace to use (host|private)")
	container_runCmd.Flags().String("cidfile", "", "Write the container ID to the file")
	container_runCmd.Flags().Int64("cpu-count", 0, "CPU count (Windows only)")
	container_runCmd.Flags().Int64("cpu-percent", 0, "CPU percent (Windows only)")
	container_runCmd.Flags().Int64("cpu-period", 0, "Limit CPU CFS (Completely Fair Scheduler) period")
	container_runCmd.Flags().Int64("cpu-quota", 0, "Limit CPU CFS (Completely Fair Scheduler) quota")
	container_runCmd.Flags().Int64("cpu-rt-period", 0, "Limit CPU real-time period in microseconds")
	container_runCmd.Flags().Int64("cpu-rt-runtime", 0, "Limit CPU real-time runtime in microseconds")
	container_runCmd.Flags().Int64P("cpu-shares", "c", 0, "CPU shares (relative weight)")
	container_runCmd.Flags().String("cpus", "", "Number of CPUs")
	container_runCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	container_runCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	container_runCmd.Flags().BoolP("detach", "d", false, "Run container in background and print container ID")
	container_runCmd.Flags().String("detach-keys", "", "Override the key sequence for detaching a container")
	container_runCmd.Flags().String("device", "", "Add a host device to the container")
	container_runCmd.Flags().String("device-cgroup-rule", "", "Add a rule to the cgroup allowed devices list")
	container_runCmd.Flags().String("device-read-bps", "", "Limit read rate (bytes per second) from a device")
	container_runCmd.Flags().String("device-read-iops", "", "Limit read rate (IO per second) from a device")
	container_runCmd.Flags().String("device-write-bps", "", "Limit write rate (bytes per second) to a device")
	container_runCmd.Flags().String("device-write-iops", "", "Limit write rate (IO per second) to a device")
	container_runCmd.Flags().Bool("disable-content-trust", true, "Skip image verification")
	container_runCmd.Flags().String("dns", "", "Set custom DNS servers")
	container_runCmd.Flags().String("dns-opt", "", "Set DNS options")
	container_runCmd.Flags().String("dns-option", "", "Set DNS options")
	container_runCmd.Flags().String("dns-search", "", "Set custom DNS search domains")
	container_runCmd.Flags().String("domainname", "", "Container NIS domain name")
	container_runCmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image")
	container_runCmd.Flags().StringP("env", "e", "", "Set environment variables")
	container_runCmd.Flags().String("env-file", "", "Read in a file of environment variables")
	container_runCmd.Flags().String("expose", "", "Expose a port or a range of ports")
	container_runCmd.Flags().String("gpus", "", "GPU devices to add to the container ('all' to pass all GPUs)")
	container_runCmd.Flags().String("group-add", "", "Add additional groups to join")
	container_runCmd.Flags().String("health-cmd", "", "Command to run to check health")
	container_runCmd.Flags().Duration("health-interval", 0, "Time between running the check (ms|s|m|h) (default 0s)")
	container_runCmd.Flags().Int("health-retries", 0, "Consecutive failures needed to report unhealthy")
	container_runCmd.Flags().Duration("health-start-period", 0, "Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)")
	container_runCmd.Flags().Duration("health-timeout", 0, "Maximum time to allow one check to run (ms|s|m|h) (default 0s)")
	container_runCmd.Flags().Bool("help", false, "Print usage")
	container_runCmd.Flags().StringP("hostname", "h", "", "Container host name")
	container_runCmd.Flags().Bool("init", false, "Run an init inside the container that forwards signals and reaps processes")
	container_runCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	container_runCmd.Flags().String("io-maxbandwidth", "", "Maximum IO bandwidth limit for the system drive (Windows only)")
	container_runCmd.Flags().Uint64("io-maxiops", 0, "Maximum IOps limit for the system drive (Windows only)")
	container_runCmd.Flags().String("ip", "", "IPv4 address (e.g., 172.30.100.104)")
	container_runCmd.Flags().String("ip6", "", "IPv6 address (e.g., 2001:db8::33)")
	container_runCmd.Flags().String("ipc", "", "IPC mode to use")
	container_runCmd.Flags().String("isolation", "", "Container isolation technology")
	container_runCmd.Flags().String("kernel-memory", "", "Kernel memory limit")
	container_runCmd.Flags().StringP("label", "l", "", "Set meta data on a container")
	container_runCmd.Flags().String("label-file", "", "Read in a line delimited file of labels")
	container_runCmd.Flags().String("link", "", "Add link to another container")
	container_runCmd.Flags().String("link-local-ip", "", "Container IPv4/IPv6 link-local addresses")
	container_runCmd.Flags().String("log-driver", "", "Logging driver for the container")
	container_runCmd.Flags().String("log-opt", "", "Log driver options")
	container_runCmd.Flags().String("mac-address", "", "Container MAC address (e.g., 92:d0:c6:0a:29:33)")
	container_runCmd.Flags().StringP("memory", "m", "", "Memory limit")
	container_runCmd.Flags().String("memory-reservation", "", "Memory soft limit")
	container_runCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	container_runCmd.Flags().Int64("memory-swappiness", 0, "Tune container memory swappiness (0 to 100)")
	container_runCmd.Flags().String("mount", "", "Attach a filesystem mount to the container")
	container_runCmd.Flags().String("name", "", "Assign a name to the container")
	container_runCmd.Flags().String("net", "", "Connect a container to a network")
	container_runCmd.Flags().String("net-alias", "", "Add network-scoped alias for the container")
	container_runCmd.Flags().String("network", "", "Connect a container to a network")
	container_runCmd.Flags().String("network-alias", "", "Add network-scoped alias for the container")
	container_runCmd.Flags().Bool("no-healthcheck", false, "Disable any container-specified HEALTHCHECK")
	container_runCmd.Flags().Bool("oom-kill-disable", false, "Disable OOM Killer")
	container_runCmd.Flags().Int("oom-score-adj", 0, "Tune host's OOM preferences (-1000 to 1000)")
	container_runCmd.Flags().String("pid", "", "PID namespace to use")
	container_runCmd.Flags().Int64("pids-limit", 0, "Tune container pids limit (set -1 for unlimited)")
	container_runCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	container_runCmd.Flags().Bool("privileged", false, "Give extended privileges to this container")
	container_runCmd.Flags().StringP("publish", "p", "", "Publish a container's port(s) to the host")
	container_runCmd.Flags().BoolP("publish-all", "P", false, "Publish all exposed ports to random ports")
	container_runCmd.Flags().String("pull", "missing", "Pull image before running (\"always\", \"missing\", \"never\")")
	container_runCmd.Flags().BoolP("quiet", "q", false, "Suppress the pull output")
	container_runCmd.Flags().Bool("read-only", false, "Mount the container's root filesystem as read only")
	container_runCmd.Flags().String("restart", "no", "Restart policy to apply when a container exits")
	container_runCmd.Flags().Bool("rm", false, "Automatically remove the container when it exits")
	container_runCmd.Flags().String("runtime", "", "Runtime to use for this container")
	container_runCmd.Flags().String("security-opt", "", "Security Options")
	container_runCmd.Flags().String("shm-size", "", "Size of /dev/shm")
	container_runCmd.Flags().Bool("sig-proxy", true, "Proxy received signals to the process")
	container_runCmd.Flags().String("stop-signal", "", "Signal to stop the container")
	container_runCmd.Flags().Int("stop-timeout", 0, "Timeout (in seconds) to stop a container")
	container_runCmd.Flags().String("storage-opt", "", "Storage driver options for the container")
	container_runCmd.Flags().String("sysctl", "", "Sysctl options")
	container_runCmd.Flags().String("tmpfs", "", "Mount a tmpfs directory")
	container_runCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	container_runCmd.Flags().String("ulimit", "", "Ulimit options")
	container_runCmd.Flags().StringP("user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
	container_runCmd.Flags().String("userns", "", "User namespace to use")
	container_runCmd.Flags().String("uts", "", "UTS namespace to use")
	container_runCmd.Flags().StringP("volume", "v", "", "Bind mount a volume")
	container_runCmd.Flags().String("volume-driver", "", "Optional volume driver for the container")
	container_runCmd.Flags().String("volumes-from", "", "Mount volumes from the specified container(s)")
	container_runCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	containerCmd.AddCommand(container_runCmd)

	carapace.Gen(container_runCmd).FlagCompletion(carapace.ActionMap{
		"attach":       carapace.ActionValues("STDERR", "STDIN", "STDOUT"),
		"blkio-weight": carapace.ActionValues("10", "100", "500", "1000"),
		"cgroupns":     carapace.ActionValues("host", "private"),
		"cidfile":      carapace.ActionFiles(),
		"cpu-shares":   carapace.ActionValues("0", "10", "100", "200", "500", "800", "1000"),
		"detach-keys":  docker.ActionDetachKeys(),
		"device":       carapace.ActionFiles(),
		"env":          env.ActionNameValues(false),
		"env-file":     carapace.ActionFiles(),
		"group-add":    os.ActionGroups(),
		"isolation":    carapace.ActionValues("default", "hyperv", "process"),
		"label-file":   carapace.ActionFiles(),
		"log-driver":   docker.ActionLogDrivers(),
		"network":      carapace.ActionValues("bridge", "container", "host", "none"),
		"pid":          carapace.ActionValues("container", "host"),
		"user":         os.ActionUsers(),
		"volume":       docker.ActionVolumes(),
		"volumes-from": carapace.Batch(
			docker.ActionContainers(),
			docker.ActionContainerIds(),
		).ToA(),
	})

	carapace.Gen(container_runCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
