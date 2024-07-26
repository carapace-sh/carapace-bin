package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueWorkCmd = &cobra.Command{
	Use:   "queue:work [--name [NAME]] [--queue [QUEUE]] [--daemon] [--once] [--stop-when-empty] [--delay [DELAY]] [--backoff [BACKOFF]] [--max-jobs [MAX-JOBS]] [--max-time [MAX-TIME]] [--force] [--memory [MEMORY]] [--sleep [SLEEP]] [--rest [REST]] [--timeout [TIMEOUT]] [--tries [TRIES]] [--] [<connection>]",
	Short: "Start processing jobs on the queue as a daemon",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueWorkCmd).Standalone()

	queueWorkCmd.Flags().String("backoff", "", "The number of seconds to wait before retrying a job that encountered an uncaught exception")
	queueWorkCmd.Flags().Bool("daemon", false, "Run the worker in daemon mode (Deprecated)")
	queueWorkCmd.Flags().String("delay", "", "The number of seconds to delay failed jobs (Deprecated)")
	queueWorkCmd.Flags().Bool("force", false, "Force the worker to run even in maintenance mode")
	queueWorkCmd.Flags().String("max-jobs", "", "The number of jobs to process before stopping")
	queueWorkCmd.Flags().String("max-time", "", "The maximum number of seconds the worker should run")
	queueWorkCmd.Flags().String("memory", "", "The memory limit in megabytes")
	queueWorkCmd.Flags().String("name", "", "The name of the worker")
	queueWorkCmd.Flags().Bool("once", false, "Only process the next job on the queue")
	queueWorkCmd.Flags().String("queue", "", "The names of the queues to work")
	queueWorkCmd.Flags().String("rest", "", "Number of seconds to rest between jobs")
	queueWorkCmd.Flags().String("sleep", "", "Number of seconds to sleep when no job is available")
	queueWorkCmd.Flags().Bool("stop-when-empty", false, "Stop when the queue is empty")
	queueWorkCmd.Flags().String("timeout", "", "The number of seconds a child process can run")
	queueWorkCmd.Flags().String("tries", "", "Number of times to attempt a job before logging it failed")
	rootCmd.AddCommand(queueWorkCmd)
}
