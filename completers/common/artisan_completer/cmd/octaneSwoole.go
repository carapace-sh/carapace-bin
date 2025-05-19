package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneSwooleCmd = &cobra.Command{
	Use:   "octane:swoole [--host [HOST]] [--port [PORT]] [--workers [WORKERS]] [--task-workers [TASK-WORKERS]] [--max-requests [MAX-REQUESTS]] [--watch] [--poll]",
	Short: "Start the Octane Swoole server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneSwooleCmd).Standalone()

	octaneSwooleCmd.Flags().String("host", "", "The IP address the server should bind to")
	octaneSwooleCmd.Flags().String("max-requests", "", "The number of requests to process before reloading the server")
	octaneSwooleCmd.Flags().Bool("poll", false, "Use file system polling while watching in order to watch files over a network")
	octaneSwooleCmd.Flags().String("port", "", "The port the server should be available on")
	octaneSwooleCmd.Flags().String("task-workers", "", "The number of task workers that should be available to handle tasks")
	octaneSwooleCmd.Flags().Bool("watch", false, "Automatically reload the server when the application is modified")
	octaneSwooleCmd.Flags().String("workers", "", "The number of workers that should be available to handle requests")
	rootCmd.AddCommand(octaneSwooleCmd)
}
