package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var queueListenCmd = &cobra.Command{
	Use:   "queue:listen [--name [NAME]] [--delay [DELAY]] [--backoff [BACKOFF]] [--force] [--memory [MEMORY]] [--queue [QUEUE]] [--sleep [SLEEP]] [--rest [REST]] [--timeout [TIMEOUT]] [--tries [TRIES]] [--] [<connection>]",
	Short: "Listen to a given queue",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(queueListenCmd).Standalone()

	queueListenCmd.Flags().String("backoff", "", "The number of seconds to wait before retrying a job that encountered an uncaught exception")
	queueListenCmd.Flags().String("delay", "", "The number of seconds to delay failed jobs (Deprecated)")
	queueListenCmd.Flags().Bool("force", false, "Force the worker to run even in maintenance mode")
	queueListenCmd.Flags().String("memory", "", "The memory limit in megabytes")
	queueListenCmd.Flags().String("name", "", "The name of the worker")
	queueListenCmd.Flags().String("queue", "", "The queue to listen on")
	queueListenCmd.Flags().String("rest", "", "Number of seconds to rest between jobs")
	queueListenCmd.Flags().String("sleep", "", "Number of seconds to sleep when no job is available")
	queueListenCmd.Flags().String("timeout", "", "The number of seconds a child process can run")
	queueListenCmd.Flags().String("tries", "", "Number of times to attempt a job before logging it failed")
	rootCmd.AddCommand(queueListenCmd)
}
