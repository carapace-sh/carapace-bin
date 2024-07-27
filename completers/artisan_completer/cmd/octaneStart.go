package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneStartCmd = &cobra.Command{
	Use:   "octane:start [--server [SERVER]] [--host [HOST]] [--port [PORT]] [--admin-port [ADMIN-PORT]] [--rpc-host [RPC-HOST]] [--rpc-port [RPC-PORT]] [--workers [WORKERS]] [--task-workers [TASK-WORKERS]] [--max-requests [MAX-REQUESTS]] [--rr-config [RR-CONFIG]] [--caddyfile [CADDYFILE]] [--https] [--http-redirect] [--watch] [--poll] [--log-level [LOG-LEVEL]]",
	Short: "Start the Octane server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneStartCmd).Standalone()

	octaneStartCmd.Flags().String("admin-port", "", "The port the admin server should be available on [FrankenPHP only]")
	octaneStartCmd.Flags().String("caddyfile", "", "The path to the FrankenPHP Caddyfile file")
	octaneStartCmd.Flags().String("host", "", "The IP address the server should bind to")
	octaneStartCmd.Flags().Bool("http-redirect", false, "Enable HTTP to HTTPS redirection (only enabled if --https is passed) [FrankenPHP only]")
	octaneStartCmd.Flags().Bool("https", false, "Enable HTTPS, HTTP/2, and HTTP/3, and automatically generate and renew certificates [FrankenPHP only]")
	octaneStartCmd.Flags().String("log-level", "", "Log messages at or above the specified log level")
	octaneStartCmd.Flags().String("max-requests", "", "The number of requests to process before reloading the server")
	octaneStartCmd.Flags().Bool("poll", false, "Use file system polling while watching in order to watch files over a network")
	octaneStartCmd.Flags().String("port", "", "The port the server should be available on [default: \"8000\"]")
	octaneStartCmd.Flags().String("rpc-host", "", "The RPC IP address the server should bind to")
	octaneStartCmd.Flags().String("rpc-port", "", "The RPC port the server should be available on")
	octaneStartCmd.Flags().String("rr-config", "", "The path to the RoadRunner .rr.yaml file")
	octaneStartCmd.Flags().String("server", "", "The server that should be used to serve the application")
	octaneStartCmd.Flags().String("task-workers", "", "The number of task workers that should be available to handle tasks")
	octaneStartCmd.Flags().Bool("watch", false, "Automatically reload the server when the application is modified")
	octaneStartCmd.Flags().String("workers", "", "The number of workers that should be available to handle requests")
	rootCmd.AddCommand(octaneStartCmd)
}
