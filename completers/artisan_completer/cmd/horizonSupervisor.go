package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonSupervisorCmd = &cobra.Command{
	Use:   "horizon:supervisor [--balance [BALANCE]] [--delay [DELAY]] [--backoff [BACKOFF]] [--max-jobs [MAX-JOBS]] [--max-time [MAX-TIME]] [--force] [--max-processes [MAX-PROCESSES]] [--min-processes [MIN-PROCESSES]] [--memory [MEMORY]] [--nice [NICE]] [--paused] [--queue [QUEUE]] [--sleep [SLEEP]] [--timeout [TIMEOUT]] [--tries [TRIES]] [--auto-scaling-strategy [AUTO-SCALING-STRATEGY]] [--balance-cooldown [BALANCE-COOLDOWN]] [--balance-max-shift [BALANCE-MAX-SHIFT]] [--workers-name [WORKERS-NAME]] [--parent-id [PARENT-ID]] [--rest [REST]] [--] <name> <connection>",
	Short: "Start a new supervisor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonSupervisorCmd).Standalone()

	horizonSupervisorCmd.Flags().String("auto-scaling-strategy", "", "If supervisor should scale by jobs or time to complete")
	horizonSupervisorCmd.Flags().String("backoff", "", "The number of seconds to wait before retrying a job that encountered an uncaught exception")
	horizonSupervisorCmd.Flags().String("balance", "", "The balancing strategy the supervisor should apply")
	horizonSupervisorCmd.Flags().String("balance-cooldown", "", "The number of seconds to wait in between auto-scaling attempts")
	horizonSupervisorCmd.Flags().String("balance-max-shift", "", "The maximum number of processes to increase or decrease per one scaling")
	horizonSupervisorCmd.Flags().String("delay", "", "The number of seconds to delay failed jobs (Deprecated)")
	horizonSupervisorCmd.Flags().Bool("force", false, "Force the worker to run even in maintenance mode")
	horizonSupervisorCmd.Flags().String("max-jobs", "", "The number of jobs to process before stopping a child process")
	horizonSupervisorCmd.Flags().String("max-processes", "", "The maximum number of total workers to start")
	horizonSupervisorCmd.Flags().String("max-time", "", "The maximum number of seconds a child process should run")
	horizonSupervisorCmd.Flags().String("memory", "", "The memory limit in megabytes")
	horizonSupervisorCmd.Flags().String("min-processes", "", "The minimum number of workers to assign per queue")
	horizonSupervisorCmd.Flags().String("nice", "", "The process priority")
	horizonSupervisorCmd.Flags().String("parent-id", "", "The parent process ID")
	horizonSupervisorCmd.Flags().Bool("paused", false, "Start the supervisor in a paused state")
	horizonSupervisorCmd.Flags().String("queue", "", "The names of the queues to work")
	horizonSupervisorCmd.Flags().String("rest", "", "Number of seconds to rest between jobs")
	horizonSupervisorCmd.Flags().String("sleep", "", "Number of seconds to sleep when no job is available")
	horizonSupervisorCmd.Flags().String("timeout", "", "The number of seconds a child process can run")
	horizonSupervisorCmd.Flags().String("tries", "", "Number of times to attempt a job before logging it failed")
	horizonSupervisorCmd.Flags().String("workers-name", "", "The name that should be assigned to the workers")
	rootCmd.AddCommand(horizonSupervisorCmd)
}
