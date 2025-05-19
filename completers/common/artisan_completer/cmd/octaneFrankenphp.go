package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var octaneFrankenphpCmd = &cobra.Command{
	Use:   "octane:frankenphp [--host [HOST]] [--port [PORT]] [--admin-host [ADMIN-HOST]] [--admin-port [ADMIN-PORT]] [--workers [WORKERS]] [--max-requests [MAX-REQUESTS]] [--caddyfile [CADDYFILE]] [--https] [--http-redirect] [--watch] [--poll] [--log-level [LOG-LEVEL]]",
	Short: "Start the Octane FrankenPHP server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(octaneFrankenphpCmd).Standalone()

	octaneFrankenphpCmd.Flags().String("admin-host", "", "The host the admin server should be available on")
	octaneFrankenphpCmd.Flags().String("admin-port", "", "The port the admin server should be available on")
	octaneFrankenphpCmd.Flags().String("caddyfile", "", "The path to the FrankenPHP Caddyfile file")
	octaneFrankenphpCmd.Flags().String("host", "", "The IP address the server should bind to")
	octaneFrankenphpCmd.Flags().Bool("http-redirect", false, "Enable HTTP to HTTPS redirection (only enabled if --https is passed)")
	octaneFrankenphpCmd.Flags().Bool("https", false, "Enable HTTPS, HTTP/2, and HTTP/3, and automatically generate and renew certificates")
	octaneFrankenphpCmd.Flags().String("log-level", "", "Log messages at or above the specified log level")
	octaneFrankenphpCmd.Flags().String("max-requests", "", "The number of requests to process before reloading the server")
	octaneFrankenphpCmd.Flags().Bool("poll", false, "Use file system polling while watching in order to watch files over a network")
	octaneFrankenphpCmd.Flags().String("port", "", "The port the server should be available on")
	octaneFrankenphpCmd.Flags().Bool("watch", false, "Automatically reload the server when the application is modified")
	octaneFrankenphpCmd.Flags().String("workers", "", "The number of workers that should be available to handle requests")
	rootCmd.AddCommand(octaneFrankenphpCmd)
}
