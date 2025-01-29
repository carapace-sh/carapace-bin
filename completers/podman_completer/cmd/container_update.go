package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var container_updateCmd = &cobra.Command{
	Use:   "update [options] CONTAINER",
	Short: "Update an existing container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(container_updateCmd).Standalone()

	container_updateCmd.Flags().String("blkio-weight", "", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.")
	container_updateCmd.Flags().StringSlice("blkio-weight-device", []string{}, "Block IO weight (relative device weight, format: `DEVICE_NAME:WEIGHT`)")
	container_updateCmd.Flags().String("cpu-period", "", "Limit the CPU CFS (Completely Fair Scheduler) period")
	container_updateCmd.Flags().String("cpu-quota", "", "Limit the CPU CFS (Completely Fair Scheduler) quota")
	container_updateCmd.Flags().String("cpu-rt-period", "", "Limit the CPU real-time period in microseconds")
	container_updateCmd.Flags().String("cpu-rt-runtime", "", "Limit the CPU real-time runtime in microseconds")
	container_updateCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	container_updateCmd.Flags().String("cpus", "", "Number of CPUs. The default is 0.000 which means no limit")
	container_updateCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	container_updateCmd.Flags().String("cpuset-mems", "", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.")
	container_updateCmd.Flags().StringSlice("device-read-bps", []string{}, "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)")
	container_updateCmd.Flags().StringSlice("device-read-iops", []string{}, "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)")
	container_updateCmd.Flags().StringSlice("device-write-bps", []string{}, "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)")
	container_updateCmd.Flags().StringSlice("device-write-iops", []string{}, "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)")
	container_updateCmd.Flags().String("health-cmd", "", "set a healthcheck command for the container ('none' disables the existing healthcheck)")
	container_updateCmd.Flags().String("health-interval", "", "set an interval for the healthcheck. (a value of disable results in no automatic timer setup) Changing this setting resets timer.")
	container_updateCmd.Flags().String("health-log-destination", "", "set the destination of the HealthCheck log. Directory path, local or events_logger (local use container state file) Warning: Changing this setting may cause the loss of previous logs!")
	container_updateCmd.Flags().String("health-max-log-count", "", "set maximum number of attempts in the HealthCheck log file. ('0' value means an infinite number of attempts in the log file)")
	container_updateCmd.Flags().String("health-max-log-size", "", "set maximum length in characters of stored HealthCheck log. ('0' value means an infinite log length)")
	container_updateCmd.Flags().String("health-on-failure", "", "action to take once the container turns unhealthy")
	container_updateCmd.Flags().String("health-retries", "", "the number of retries allowed before a healthcheck is considered to be unhealthy")
	container_updateCmd.Flags().String("health-start-period", "", "the initialization time needed for a container to bootstrap")
	container_updateCmd.Flags().String("health-startup-cmd", "", "Set a startup healthcheck command for the container")
	container_updateCmd.Flags().String("health-startup-interval", "", "Set an interval for the startup healthcheck. Changing this setting resets the timer, depending on the state of the container.")
	container_updateCmd.Flags().String("health-startup-retries", "", "Set the maximum number of retries before the startup healthcheck will restart the container")
	container_updateCmd.Flags().String("health-startup-success", "", "Set the number of consecutive successes before the startup healthcheck is marked as successful and the normal healthcheck begins (0 indicates any success will start the regular healthcheck)")
	container_updateCmd.Flags().String("health-startup-timeout", "", "Set the maximum amount of time that the startup healthcheck may take before it is considered failed")
	container_updateCmd.Flags().String("health-timeout", "", "the maximum time allowed to complete the healthcheck before an interval is considered failed")
	container_updateCmd.Flags().StringP("memory", "m", "", "Memory limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	container_updateCmd.Flags().String("memory-reservation", "", "Memory soft limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	container_updateCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	container_updateCmd.Flags().String("memory-swappiness", "", "Tune container memory swappiness (0 to 100, or -1 for system default)")
	container_updateCmd.Flags().Bool("no-healthcheck", false, "Disable healthchecks on container")
	container_updateCmd.Flags().String("pids-limit", "", "Tune container pids limit (set -1 for unlimited)")
	container_updateCmd.Flags().String("restart", "", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")")
	containerCmd.AddCommand(container_updateCmd)
}
