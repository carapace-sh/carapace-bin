package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "update [options] CONTAINER",
	Short: "Update an existing container",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(updateCmd).Standalone()

	updateCmd.Flags().String("blkio-weight", "", "Block IO weight (relative weight) accepts a weight value between 10 and 1000.")
	updateCmd.Flags().StringSlice("blkio-weight-device", []string{}, "Block IO weight (relative device weight, format: `DEVICE_NAME:WEIGHT`)")
	updateCmd.Flags().String("cpu-period", "", "Limit the CPU CFS (Completely Fair Scheduler) period")
	updateCmd.Flags().String("cpu-quota", "", "Limit the CPU CFS (Completely Fair Scheduler) quota")
	updateCmd.Flags().String("cpu-rt-period", "", "Limit the CPU real-time period in microseconds")
	updateCmd.Flags().String("cpu-rt-runtime", "", "Limit the CPU real-time runtime in microseconds")
	updateCmd.Flags().StringP("cpu-shares", "c", "", "CPU shares (relative weight)")
	updateCmd.Flags().String("cpus", "", "Number of CPUs. The default is 0.000 which means no limit")
	updateCmd.Flags().String("cpuset-cpus", "", "CPUs in which to allow execution (0-3, 0,1)")
	updateCmd.Flags().String("cpuset-mems", "", "Memory nodes (MEMs) in which to allow execution (0-3, 0,1). Only effective on NUMA systems.")
	updateCmd.Flags().StringSlice("device-read-bps", []string{}, "Limit read rate (bytes per second) from a device (e.g. --device-read-bps=/dev/sda:1mb)")
	updateCmd.Flags().StringSlice("device-read-iops", []string{}, "Limit read rate (IO per second) from a device (e.g. --device-read-iops=/dev/sda:1000)")
	updateCmd.Flags().StringSlice("device-write-bps", []string{}, "Limit write rate (bytes per second) to a device (e.g. --device-write-bps=/dev/sda:1mb)")
	updateCmd.Flags().StringSlice("device-write-iops", []string{}, "Limit write rate (IO per second) to a device (e.g. --device-write-iops=/dev/sda:1000)")
	updateCmd.Flags().String("health-cmd", "", "set a healthcheck command for the container ('none' disables the existing healthcheck)")
	updateCmd.Flags().String("health-interval", "", "set an interval for the healthcheck. (a value of disable results in no automatic timer setup) Changing this setting resets timer.")
	updateCmd.Flags().String("health-log-destination", "", "set the destination of the HealthCheck log. Directory path, local or events_logger (local use container state file) Warning: Changing this setting may cause the loss of previous logs!")
	updateCmd.Flags().String("health-max-log-count", "", "set maximum number of attempts in the HealthCheck log file. ('0' value means an infinite number of attempts in the log file)")
	updateCmd.Flags().String("health-max-log-size", "", "set maximum length in characters of stored HealthCheck log. ('0' value means an infinite log length)")
	updateCmd.Flags().String("health-on-failure", "", "action to take once the container turns unhealthy")
	updateCmd.Flags().String("health-retries", "", "the number of retries allowed before a healthcheck is considered to be unhealthy")
	updateCmd.Flags().String("health-start-period", "", "the initialization time needed for a container to bootstrap")
	updateCmd.Flags().String("health-startup-cmd", "", "Set a startup healthcheck command for the container")
	updateCmd.Flags().String("health-startup-interval", "", "Set an interval for the startup healthcheck. Changing this setting resets the timer, depending on the state of the container.")
	updateCmd.Flags().String("health-startup-retries", "", "Set the maximum number of retries before the startup healthcheck will restart the container")
	updateCmd.Flags().String("health-startup-success", "", "Set the number of consecutive successes before the startup healthcheck is marked as successful and the normal healthcheck begins (0 indicates any success will start the regular healthcheck)")
	updateCmd.Flags().String("health-startup-timeout", "", "Set the maximum amount of time that the startup healthcheck may take before it is considered failed")
	updateCmd.Flags().String("health-timeout", "", "the maximum time allowed to complete the healthcheck before an interval is considered failed")
	updateCmd.Flags().StringP("memory", "m", "", "Memory limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	updateCmd.Flags().String("memory-reservation", "", "Memory soft limit (format: `<number>[<unit>]`, where unit = b (bytes), k (kibibytes), m (mebibytes), or g (gibibytes))")
	updateCmd.Flags().String("memory-swap", "", "Swap limit equal to memory plus swap: '-1' to enable unlimited swap")
	updateCmd.Flags().String("memory-swappiness", "", "Tune container memory swappiness (0 to 100, or -1 for system default)")
	updateCmd.Flags().Bool("no-healthcheck", false, "Disable healthchecks on container")
	updateCmd.Flags().String("pids-limit", "", "Tune container pids limit (set -1 for unlimited)")
	updateCmd.Flags().String("restart", "", "Restart policy to apply when a container exits (\"always\"|\"no\"|\"never\"|\"on-failure\"|\"unless-stopped\")")
	rootCmd.AddCommand(updateCmd)
}
