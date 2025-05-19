package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonWorkCmd = &cobra.Command{
	Use:   "horizon:work [--name [NAME]] [--delay [DELAY]] [--backoff [BACKOFF]] [--max-jobs [MAX-JOBS]] [--max-time [MAX-TIME]] [--daemon] [--force] [--memory [MEMORY]] [--once] [--stop-when-empty] [--queue [QUEUE]] [--sleep [SLEEP]] [--rest [REST]] [--supervisor [SUPERVISOR]] [--timeout [TIMEOUT]] [--tries [TRIES]] [--] [<connection>]",
	Short: "Start processing jobs on the queue as a daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonWorkCmd).Standalone()

	horizonWorkCmd.Flags().String("backoff", "", "The number of seconds to wait before retrying a job that encountered an uncaught exception")
	horizonWorkCmd.Flags().Bool("daemon", false, "Run the worker in daemon mode (Deprecated)")
	horizonWorkCmd.Flags().String("delay", "", "The number of seconds to delay failed jobs (Deprecated)")
	horizonWorkCmd.Flags().Bool("force", false, "Force the worker to run even in maintenance mode")
	horizonWorkCmd.Flags().String("max-jobs", "", "The number of jobs to process before stopping")
	horizonWorkCmd.Flags().String("max-time", "", "The maximum number of seconds the worker should run")
	horizonWorkCmd.Flags().String("memory", "", "The memory limit in megabytes")
	horizonWorkCmd.Flags().String("name", "", "The name of the worker")
	horizonWorkCmd.Flags().Bool("once", false, "Only process the next job on the queue")
	horizonWorkCmd.Flags().String("queue", "", "The names of the queues to work")
	horizonWorkCmd.Flags().String("rest", "", "Number of seconds to rest between jobs")
	horizonWorkCmd.Flags().String("sleep", "", "Number of seconds to sleep when no job is available")
	horizonWorkCmd.Flags().Bool("stop-when-empty", false, "Stop when the queue is empty")
	horizonWorkCmd.Flags().String("supervisor", "", "The name of the supervisor the worker belongs to")
	horizonWorkCmd.Flags().String("timeout", "", "The number of seconds a child process can run")
	horizonWorkCmd.Flags().String("tries", "", "Number of times to attempt a job before logging it failed")
	rootCmd.AddCommand(horizonWorkCmd)
}
