package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/net"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve [--host [HOST]] [--port [PORT]] [--tries [TRIES]] [--no-reload]",
	Short: "Serve the application on the PHP development server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serveCmd).Standalone()

	serveCmd.Flags().String("host", "", "The host address to serve the application on")
	serveCmd.Flags().Bool("no-reload", false, "Do not reload the development server on .env file changes")
	serveCmd.Flags().String("port", "", "The port to serve the application on")
	serveCmd.Flags().String("tries", "", "The max number of ports to attempt to serve from")
	rootCmd.AddCommand(serveCmd)

	carapace.Gen(serveCmd).FlagCompletion(carapace.ActionMap{
		"host": carapace.ActionValues("localhost", "127.0.0.1"),
		"port": net.ActionPorts(),
	})
}
