package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/env"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/docker"
	"github.com/spf13/cobra"
)

var container_createCmd = &cobra.Command{
	Use:   "create [OPTIONS] IMAGE [COMMAND] [ARG...]",
	Short: "Create a new container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_createCmd).Standalone()

	container_createCmd.Flags().String("add-host", "", "Add a custom host-to-IP mapping (host:ip)")
	container_createCmd.Flags().StringP("attach", "a", "", "Attach to STDIN, STDOUT or STDERR")
	container_createCmd.Flags().Uint16("blkio-weight", 0, "Block IO (relative weight), between 10 and 1000, or 0 to disable (default 0)")
	container_createCmd.Flags().String("blkio-weight-device", "", "Block IO weight (relative device weight)")
	container_createCmd.Flags().String("cap-add", "", "Add Linux capabilities")
	container_createCmd.Flags().String("cap-drop", "", "Drop Linux capabilities")
	container_createCmd.Flags().String("cgroup-parent", "", "Optional parent cgroup for the container")
	container_createCmd.Flags().String("cgroupns", "", "Cgroup namespace to use (host|private)")
	container_createCmd.Flags().String("cidfile", "", "Write the container ID to the file")
	container_createCmd.Flags().Int64("cpu-count", 0, "CPU count (Windows only)")
	container_createCmd.Flags().Int64("cpu-percent", 0, "CPU percent (Windows only)")
	container_createCmd.Flags().Int64("cpu-period", 0, "Limit CPU CFS (Completely Fair Scheduler) period")
	container_createCmd.Flags().Int64("cpu-quota", 0, "Limit CPU CFS (Completely Fair Scheduler) quota")
	container_createCmd.Flags().Int64("cpu-rt-period", 0, "Limit CPU real-time period in microseconds")
	container_createCmd.Flags().Int64("cpu-rt-runtime", 0, "Limit CPU real-time runtime in microseconds")
	container_createCmd.Flags().Int64P("cpu-shares", "c", 0, "CPU shares (relative weight)")
	container_createCmd.Flags().String("cpus", "", "Number of CPUs")
	container_createCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	container_createCmd.Flags().String("cpuset-mems", "", "MEMs in which to allow execution (0-3, 0,1)")
	container_createCmd.Flags().String("device", "", "Add a host device to the container")
	container_createCmd.Flags().String("device-cgroup-rule", "", "Add a rule to the cgroup allowed devices list")
	container_createCmd.Flags().String("device-read-bps", "", "Limit read rate (bytes per second) from a device")
	container_createCmd.Flags().String("device-read-iops", "", "Limit read rate (IO per second) from a device")
	container_createCmd.Flags().String("device-write-bps", "", "Limit write rate (bytes per second) to a device")
	container_createCmd.Flags().String("device-write-iops", "", "Limit write rate (IO per second) to a device")
	container_createCmd.Flags().Bool("disable-content-trust", true, "Skip image verification")
	container_createCmd.Flags().String("dns", "", "Set custom DNS servers")
	container_createCmd.Flags().String("dns-opt", "", "Set DNS options")
	container_createCmd.Flags().String("dns-option", "", "Set DNS options")
	container_createCmd.Flags().String("dns-search", "", "Set custom DNS search domains")
	container_createCmd.Flags().String("domainname", "", "Container NIS domain name")
	container_createCmd.Flags().String("entrypoint", "", "Overwrite the default ENTRYPOINT of the image")
	container_createCmd.Flags().StringP("env", "e", "", "Set environment variables")
	container_createCmd.Flags().String("env-file", "", "Read in a file of environment variables")
	container_createCmd.Flags().String("expose", "", "Expose a port or a range of ports")
	container_createCmd.Flags().String("gpus", "", "GPU devices to add to the container ('all' to pass all GPUs)")
	container_createCmd.Flags().String("group-add", "", "Add additional groups to join")
	container_createCmd.Flags().String("health-cmd", "", "Command to run to check health")
	container_createCmd.Flags().Duration("health-interval", 0, "Time between running the check (ms|s|m|h) (default 0s)")
	container_createCmd.Flags().Int("health-retries", 0, "Consecutive failures needed to report unhealthy")
	container_createCmd.Flags().Duration("health-start-period", 0, "Start period for the container to initialize before starting health-retries countdown (ms|s|m|h) (default 0s)")
	container_createCmd.Flags().Duration("health-timeout", 0, "Maximum time to allow one check to run (ms|s|m|h) (default 0s)")
	container_createCmd.Flags().Bool("help", false, "Print usage")
	container_createCmd.Flags().StringP("hostname", "h", "", "Container host name")
	container_createCmd.Flags().Bool("init", false, "Run an init inside the container that forwards signals and reaps processes")
	container_createCmd.Flags().BoolP("interactive", "i", false, "Keep STDIN open even if not attached")
	container_createCmd.Flags().String("io-maxbandwidth", "", "Maximum IO bandwidth limit for the system drive (Windows only)")
	container_createCmd.Flags().Uint64("io-maxiops", 0, "Maximum IOps limit for the system drive (Windows only)")
	container_createCmd.Flags().String("ip", "", "IPv4 address (e.g., 172.30.100.104)")
	container_createCmd.Flags().String("ip6", "", "IPv6 address (e.g., 2001:db8::33)")
	container_createCmd.Flags().String("ipc", "", "IPC mode to use")
	container_createCmd.Flags().String("isolation", "", "Container isolation technology")
	container_createCmd.Flags().String("kernel-memory", "", "Kernel memory limit")
	container_createCmd.Flags().StringP("label", "l", "", "Set meta data on a container")
	container_createCmd.Flags().String("label-file", "", "Read in a line delimited file of labels")
	container_createCmd.Flags().String("link", "", "Add link to another container")
	container_createCmd.Flags().String("link-local-ip", "", "Container IPv4/IPv6 link-local addresses")
	container_createCmd.Flags().String("log-driver", "", "Logging driver for the container")
	container_createCmd.Flags().String("log-opt", "", "Log driver options")
	container_createCmd.Flags().String("mac-address", "", "Container MAC address (e.g., 92:d0:c6:0a:29:33)")
	container_createCmd.Flags().StringP("memory", "m", "", "Memory limit")
	container_createCmd.Flags().String("memory-reservation", "", "Memory soft limit")
	container_createCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	container_createCmd.Flags().Int64("memory-swappiness", 0, "Tune container memory swappiness (0 to 100)")
	container_createCmd.Flags().String("mount", "", "Attach a filesystem mount to the container")
	container_createCmd.Flags().String("name", "", "Assign a name to the container")
	container_createCmd.Flags().String("net", "", "Connect a container to a network")
	container_createCmd.Flags().String("net-alias", "", "Add network-scoped alias for the container")
	container_createCmd.Flags().String("network", "", "Connect a container to a network")
	container_createCmd.Flags().String("network-alias", "", "Add network-scoped alias for the container")
	container_createCmd.Flags().Bool("no-healthcheck", false, "Disable any container-specified HEALTHCHECK")
	container_createCmd.Flags().Bool("oom-kill-disable", false, "Disable OOM Killer")
	container_createCmd.Flags().Int("oom-score-adj", 0, "Tune host's OOM preferences (-1000 to 1000)")
	container_createCmd.Flags().String("pid", "", "PID namespace to use")
	container_createCmd.Flags().Int64("pids-limit", 0, "Tune container pids limit (set -1 for unlimited)")
	container_createCmd.Flags().String("platform", "", "Set platform if server is multi-platform capable")
	container_createCmd.Flags().Bool("privileged", false, "Give extended privileges to this container")
	container_createCmd.Flags().StringP("publish", "p", "", "Publish a container's port(s) to the host")
	container_createCmd.Flags().BoolP("publish-all", "P", false, "Publish all exposed ports to random ports")
	container_createCmd.Flags().String("pull", "missing", "Pull image before creating (\"always\", \"|missing\", \"never\")")
	container_createCmd.Flags().BoolP("quiet", "q", false, "Suppress the pull output")
	container_createCmd.Flags().Bool("read-only", false, "Mount the container's root filesystem as read only")
	container_createCmd.Flags().String("restart", "no", "Restart policy to apply when a container exits")
	container_createCmd.Flags().Bool("rm", false, "Automatically remove the container when it exits")
	container_createCmd.Flags().String("runtime", "", "Runtime to use for this container")
	container_createCmd.Flags().String("security-opt", "", "Security Options")
	container_createCmd.Flags().String("shm-size", "", "Size of /dev/shm")
	container_createCmd.Flags().String("stop-signal", "", "Signal to stop the container")
	container_createCmd.Flags().Int("stop-timeout", 0, "Timeout (in seconds) to stop a container")
	container_createCmd.Flags().String("storage-opt", "", "Storage driver options for the container")
	container_createCmd.Flags().String("sysctl", "", "Sysctl options")
	container_createCmd.Flags().String("tmpfs", "", "Mount a tmpfs directory")
	container_createCmd.Flags().BoolP("tty", "t", false, "Allocate a pseudo-TTY")
	container_createCmd.Flags().String("ulimit", "", "Ulimit options")
	container_createCmd.Flags().StringP("user", "u", "", "Username or UID (format: <name|uid>[:<group|gid>])")
	container_createCmd.Flags().String("userns", "", "User namespace to use")
	container_createCmd.Flags().String("uts", "", "UTS namespace to use")
	container_createCmd.Flags().StringP("volume", "v", "", "Bind mount a volume")
	container_createCmd.Flags().String("volume-driver", "", "Optional volume driver for the container")
	container_createCmd.Flags().String("volumes-from", "", "Mount volumes from the specified container(s)")
	container_createCmd.Flags().StringP("workdir", "w", "", "Working directory inside the container")
	containerCmd.AddCommand(container_createCmd)

	carapace.Gen(container_createCmd).FlagCompletion(carapace.ActionMap{
		"attach":       carapace.ActionValues("STDERR", "STDIN", "STDOUT"),
		"blkio-weight": carapace.ActionValues("10", "100", "500", "1000"),
		"cgroupns":     carapace.ActionValues("host", "private"),
		"cidfile":      carapace.ActionFiles(),
		"cpu-shares":   carapace.ActionValues("0", "10", "100", "200", "500", "800", "1000"),
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
	})

	carapace.Gen(container_createCmd).PositionalCompletion(
		docker.ActionRepositoryTags(),
	)
}
